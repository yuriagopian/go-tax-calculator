package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IF_Its_Gets_An_Error_If_ID_Is_Black(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
	err := order.Validate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
