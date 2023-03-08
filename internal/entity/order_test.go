package entity

import "testing"

func Test_IF_Its_Gets_An_Error_If_ID_Is_Black(t *testing.T) {
	order := Order{}
	err := order.Validate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
