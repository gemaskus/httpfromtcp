package request

import (
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestfromReader(reader io.Reader) (*Request, error) {

	requestString, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Could not read the request from the io.reader: %v", err)
	}

	messageLines := strings.Split(string(requestString), "\r\n")

	requestLineElements := strings.Split(messageLines[0], " ")

	if !isAllUpperCase(requestLineElements[0]) {
		return &Request{}, fmt.Errorf("Method is not all uppercase: %s", requestLineElements[0])
	}

	httpVersion := strings.Split(requestLineElements[2], "/")

	if httpVersion[1] != "1.1" {
		return &Request{}, fmt.Errorf("Wrong HTTP version: %s", httpVersion[1])
	}

	return &Request{
		RequestLine{
			HttpVersion:   httpVersion[1],
			RequestTarget: requestLineElements[1],
			Method:        requestLineElements[0],
		},
	}, nil
}

func isAllUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
