package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const apiBaseURL = "https://api.pushmark.app"

// PushRequest represents the JSON payload for the API request
type PushRequest struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Color codes for terminal output
const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func main() {
	app := &cli.App{
		Name:    "pushmark",
		Usage:   "Simple push notification CLI tool",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name: "pushmark",
			},
		},
		Description: "Send push notifications via pushmark.app API",
		ArgsUsage:   "<channelHash> <message>",
		Action:      pushAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Value:   "info",
				Usage:   "notification type (info, log, warning, success, error)",
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.NArg() != 2 {
				return fmt.Errorf("requires exactly 2 arguments: <channelHash> <message>")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s%v%s\n", colorRed, err, colorReset)
		os.Exit(1)
	}
}

func pushAction(ctx *cli.Context) error {
	channelHash := ctx.Args().Get(0)
	message := ctx.Args().Get(1)
	notificationType := ctx.String("type")

	// Validate notification type
	validTypes := []string{"info", "log", "warning", "success", "error"}
	if !isValidType(notificationType, validTypes) {
		return fmt.Errorf("invalid notification type '%s'. Valid types are: %v", notificationType, validTypes)
	}

	if err := sendPush(channelHash, message, notificationType); err != nil {
		return fmt.Errorf("failed to send push: %w", err)
	}

	fmt.Printf("%sPush sent successfully!%s\n", colorGreen, colorReset)
	return nil
}

func isValidType(notificationType string, validTypes []string) bool {
	for _, validType := range validTypes {
		if notificationType == validType {
			return true
		}
	}
	return false
}

func sendPush(channelHash, message, notificationType string) error {
	// Prepare request payload
	payload := PushRequest{
		Message: message,
		Type:    notificationType,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	// Build the full API URL with channel hash
	fullURL := fmt.Sprintf("%s/%s", apiBaseURL, channelHash)

	// Create HTTP request
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "pushmark-cli/1.0")

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error - Status: %d", resp.StatusCode)
	}

	return nil
}
