package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	gpt "github.com/sashabaranov/go-gpt3"
)

const (
	maxTokens = 4096
	user      = "github.com/clfs/chatgpt"
)

type client struct {
	g *gpt.Client
}

func newClient(key string) *client {
	return &client{gpt.NewClient(key)}
}

func (c *client) handle(ctx context.Context, prompt string) (string, error) {
	req := gpt.CompletionRequest{
		User:      user,
		Model:     gpt.GPT3TextDavinci003,
		MaxTokens: maxTokens - len(prompt)/3, // usually works
		Prompt:    prompt,
	}

	resp, err := c.g.CreateCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed req: %w", err)
	}

	return strings.TrimSpace(resp.Choices[0].Text), nil
}

func getKey() string {
	key := os.Getenv("GPT3_KEY")
	if key == "" {
		log.Fatal("GPT3_KEY empty or unset")
	}

	return key
}

func main() {
	var (
		red   = color.New(color.FgRed).SprintFunc()
		green = color.New(color.FgGreen).SprintFunc()
	)

	client := newClient(getKey())
	ctx := context.Background()

	fmt.Println(`type "exit", "quit", or "q" to exit; hit enter to submit`)
	fmt.Printf("%s ", green(">"))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		switch input {
		case "exit", "quit", "q":
			os.Exit(0)
		}

		resp, err := client.handle(ctx, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s\n", red(">"), resp)
		fmt.Printf("%s ", green(">"))
	}
}
