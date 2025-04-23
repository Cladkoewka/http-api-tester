package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "http-tester",
		Usage: "A simple HTTP testing CLI tool",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Send a GET request",
				Action: func(c *cli.Context) error {
					url := c.Args().Get(0)
					if url == "" {
						return fmt.Errorf("URL is required")
					}
					resp, err := http.Get(url)
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Printf("Status: %s\n\n%s\n", resp.Status, body)
					return nil
				},
			},
			{
				Name:  "post",
				Usage: "Send a POST request with JSON body",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "data",
						Usage: "Raw JSON data string",
					},
					&cli.StringFlag{
						Name:  "data-file",
						Usage: "Path to file containing JSON data",
					},
				},
				Action: func(c *cli.Context) error {
					url := c.Args().Get(0)
					if url == "" {
						return fmt.Errorf("URL is required")
					}
			
					var jsonBody []byte
					var err error
			
					if dataFile := c.String("data-file"); dataFile != "" {
						jsonBody, err = os.ReadFile(dataFile)
						if err != nil {
							return fmt.Errorf("could not read file: %w", err)
						}
					} else if data := c.String("data"); data != "" {
						jsonBody = []byte(data)
					} else {
						return fmt.Errorf("either --data or --data-file must be set")
					}
			
					resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Printf("Status: %s\n\n%s\n", resp.Status, body)
					return nil
				},
			},
			{
				Name:  "put",
				Usage: "Send a PUT request with JSON body",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "data",
						Usage: "Raw JSON data string",
					},
					&cli.StringFlag{
						Name:  "data-file",
						Usage: "Path to file containing JSON data",
					},
				},
				Action: func(c *cli.Context) error {
					url := c.Args().Get(0)
					if url == "" {
						return fmt.Errorf("URL is required")
					}
			
					var jsonBody []byte
					var err error
			
					if dataFile := c.String("data-file"); dataFile != "" {
						jsonBody, err = os.ReadFile(dataFile)
						if err != nil {
							return fmt.Errorf("could not read file: %w", err)
						}
					} else if data := c.String("data"); data != "" {
						jsonBody = []byte(data)
					} else {
						return fmt.Errorf("either --data or --data-file must be set")
					}
			
					req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
					if err != nil {
						return err
					}
					req.Header.Set("Content-Type", "application/json")
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Printf("Status: %s\n\n%s\n", resp.Status, body)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Send a DELETE request",
				Action: func(c *cli.Context) error {
					url := c.Args().Get(0)
					if url == "" {
						return fmt.Errorf("URL is required")
					}
					req, err := http.NewRequest("DELETE", url, nil)
					if err != nil {
						return err
					}
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Printf("Status: %s\n\n%s\n", resp.Status, body)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
