package main

import (
	"context"
	"log"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func text2mp3(ctx context.Context, input_filename, output_filename string) {
	data, err := os.ReadFile(input_filename)
	if err != nil {
		log.Fatal(err)
	}
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: string(data[:])},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(output_filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()
	inputfile_name := "hello.txt"
	output_filename := "hello.mp3"
	text2mp3(ctx, inputfile_name, output_filename)
}
