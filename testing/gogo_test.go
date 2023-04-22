package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com:"
	if !IsValidEmail(data) {
		t.Errorf("IsValidEmail(%s) = false, want true", data)
	}
}

func TestIsValidEmail2(t *testing.T) {
	table := []struct {
		data string
		want bool
	}{
		{"email@example.com", true},
		{"email@example", false},
	}

	for _, row := range table {
		if IsValidEmail(row.data) != row.want {
			t.Errorf("IsValidEmail(%s) = %t, want %t", row.data, IsValidEmail(row.data), row.want)
		}
	}
}
