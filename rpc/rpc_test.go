package rpc_test

import (
	"learninglsp/rpc"
	"testing"
)

type EncondingExample struct {
  Testing bool
}

func TestEncode(t *testing.T) {
  expect := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

  actual := rpc.EncondeMessage(EncondingExample{ Testing: true })

  if expect != actual {
    t.Fatalf("Expected: %s, Actual: %s", expect, actual)
  }
}

func TestDecode(t *testing.T) {
  method, content, err := rpc.DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"))
  contentLength := len(content)

  if err != nil {
    t.Fatal(err)
  }

  if contentLength != 15 {
    t.Fatalf("Expect 15, Recievied: %d", contentLength)
  }

  if method != "hi" {
    t.Fatalf("Expect 'hi', recievied: %s", method)
  }
}
