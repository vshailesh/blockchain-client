package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		inputFile = flag.String("file", "", "Path to input JSON file")
		url       = flag.String("url", "https://polygon-rpc.com", "RPC endpoint URL")
		verbose   = flag.Bool("v", false, "Verbose output")
		// method    = flag.String("methodToCall", "", "Options : eth_blockNumber, eth_getBlockByNumber")
	)
	flag.Parse()

	if *verbose {
		log.Println("Starting RPC client")
		log.Printf("Target URL: %s", *url)
	}

	payload, err := getPayload(*inputFile)
	if err != nil {
		log.Fatalf("Error getting payload: %v", err)
	}

	if *verbose {
		log.Printf("Using payload: %+v", payload)
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", *url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-200 response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response ResponsePayload
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if response.Error != nil {
		log.Fatalf("RPC Error [%d]: %s", response.Error.Code, response.Error.Message)
	}

	if *verbose {
		fmt.Printf("Full response:\n%s\n", body)
	} else {
		fmt.Printf("Result: %s\n", string(response.Result))
	}
}

func getPayload(inputFile string) (RequestPayload, error) {
	var payload RequestPayload

	switch {
	case inputFile != "":
		file, err := os.Open(inputFile)
		if err != nil {
			return payload, fmt.Errorf("error opening file: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&payload); err != nil {
			return payload, fmt.Errorf("error decoding JSON from file: %w", err)
		}

	case isInputFromPipe():
		if err := json.NewDecoder(os.Stdin).Decode(&payload); err != nil {
			return payload, fmt.Errorf("error decoding JSON from stdin: %w", err)
		}

	default:
		payload = RequestPayload{
			Jsonrpc: "2.0",
			Method:  "eth_blockNumber",
			ID:      1,
		}
	}

	return payload, nil
}

func isInputFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}