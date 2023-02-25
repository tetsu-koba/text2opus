package main

import (
	"context"
	"fmt"
	"log"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func listVoices(ctx context.Context, languageCode string) {
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := texttospeechpb.ListVoicesRequest{ LanguageCode: languageCode }

	resp, err := client.ListVoices(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range resp.Voices {
		fmt.Println(x)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage:%s languageCode (ex. 'en-US', '' means all)\n", os.Args[0])
		os.Exit(1)
	}
	ctx := context.Background()
	languageCode := os.Args[1]
	listVoices(ctx, languageCode)
}
