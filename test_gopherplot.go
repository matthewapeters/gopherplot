package test_gopherplot

import (
	"testing"

	gp "github.com/matthewapeters/gopherplot"
)

func TestServer(t *testing.T) {
	go gp.Server()
}
