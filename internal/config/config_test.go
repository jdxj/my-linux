package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Printf("TG: %+v\n", TG)
	fmt.Printf("Monitor: %+v\n", Monitor)
}
