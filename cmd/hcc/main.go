package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/masa23/hugo-chatgpt-converter/config"
	"github.com/sashabaranov/go-openai"
)

var (
	version = ""
	conf    *config.Config
)

func main() {
	var err error
	var confPath string
	var inputPath string
	var outputPath string
	var showVersion bool
	flag.StringVar(&confPath, "config", "config.yaml", "path to config file")
	flag.StringVar(&inputPath, "input", "", "path to input file default is stdin")
	flag.StringVar(&outputPath, "output", "", "path to output file default is stdout")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	conf, err = config.Load(confPath)
	if err != nil {
		log.Fatal(err)
	}

	var body []byte
	// inputPathが指定されている場合はファイルから読み込む
	if inputPath != "" {
		body, err = os.ReadFile(inputPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		body, err = io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			log.Fatal(err)
		}
	}

	var output io.WriteCloser
	// outputPathが指定されている場合はファイルに書き込む
	if outputPath != "" {
		output, err = os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		defer output.Close()
	} else {
		// outputPathが指定されていない場合は標準出力に書き込む
		output = os.Stdout
	}

	client := openai.NewClient(conf.OpenAI.APIToken)
	message := fmt.Sprintf("%s\n\n%s\n", conf.Prompt, string(body))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: conf.OpenAI.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("failed to create chat completion: %v", err)
	}

	_, err = output.Write([]byte(resp.Choices[0].Message.Content))
	if err != nil {
		log.Fatalf("failed to write output: %v", err)
	}
}
