package main

import (
	"context"
	"fmt"
	"log"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func text2opus(ctx context.Context, input_filename, output_filename, name string) {
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
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: string(data)},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			Name:         name,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_OGG_OPUS,
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

func listVoices(ctx context.Context, languageCode string) {
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := texttospeechpb.ListVoicesRequest{LanguageCode: languageCode}

	resp, err := client.ListVoices(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range resp.Voices {
		fmt.Println(x)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage:\n")
		log.Printf("  %s input_text_file output_opus_file [voice_name]\n", os.Args[0])
		log.Printf("  %s -l [languageCode (ex. 'en-US')]\n", os.Args[0])

		os.Exit(1)
	}
	ctx := context.Background()
	if os.Args[1] == "-l" {
		languageCode := ""
		if len(os.Args) > 2 {
			languageCode = os.Args[2]
		}
		listVoices(ctx, languageCode)
		return
	}
	inputfile_name := os.Args[1]
	output_filename := os.Args[2]
	name := ""
	if len(os.Args) >= 4 {
		name = os.Args[3]
	}
	text2opus(ctx, inputfile_name, output_filename, name)
}
