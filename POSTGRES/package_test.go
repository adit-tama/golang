package main

import (
	service "API/service/video"
	"testing"
)

func TestSum(t *testing.T) {
	result := service.Hellow("")

	if result != "Hey you there" {
		t.Error(result)
	}
}
