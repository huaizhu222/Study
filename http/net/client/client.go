package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	body, _ := json.Marshal(map[string]string{"key": "val"})

	resp, _ := http.Post("127.0.0.1:8091", "application/json", bytes.NewReader(body))
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp: %s", respBody)
}
