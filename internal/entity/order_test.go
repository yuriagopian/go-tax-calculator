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

func Test_IF_Its_Gets_An_Error_If_Price_Is_Black_Or_Less_Than_Zero(t *testing.T) {
	order := Order{ID: "123", Price: 0}
	assert.Error(t, order.Validate(), "invalid price")

}
func Test_IF_Its_Gets_An_Error_If_Tax_Is_Black_Or_Less_Than_Zero(t *testing.T) {
	order := Order{ID: "123", Price: 0}
	assert.Error(t, order.Validate(), "invalid tax")

}
