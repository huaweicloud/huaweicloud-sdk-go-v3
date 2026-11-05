package sm3

import (
	"encoding/hex"
	"testing"
)

func TestSm3Sum(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"ascii_abc", []byte("abc"), "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0"},
		{"512bit_abcd_repeat", append(append([]byte("abcdabcdabcdabcdabcdabcdabcdabcd"), []byte("abcdabcdabcdabcdabcdabcdabcdabcd")...)), "debe9ff92275b8a138604889c18e5a4d6fdb70e5387e5765293dcba39c0c5732"},
	}

	for _, tt := range tests {
		got := hex.EncodeToString(Sm3Sum(tt.input))
		if got != tt.expected {
			t.Errorf("Sm3Sum(%s) = %s, want %s", tt.name, got, tt.expected)
		}
	}
}

func TestSm3New(t *testing.T) {
	h := New()
	h.Write([]byte("abc"))
	got := hex.EncodeToString(h.Sum(nil))
	expected := "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0"
	if got != expected {
		t.Errorf("New().Write(abc).Sum() = %s, want %s", got, expected)
	}
}

func TestSm3Reset(t *testing.T) {
	h := New()
	h.Write([]byte("abc"))
	h.Reset()
	h.Write([]byte("abc"))
	got := hex.EncodeToString(h.Sum(nil))
	expected := "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0"
	if got != expected {
		t.Errorf("After Reset: got %s, want %s", got, expected)
	}
}

func TestSm3Size(t *testing.T) {
	h := New()
	if h.Size() != 32 {
		t.Errorf("Size() = %d, want 32", h.Size())
	}
}

func TestSm3BlockSize(t *testing.T) {
	h := New()
	if h.BlockSize() != 64 {
		t.Errorf("BlockSize() = %d, want 64", h.BlockSize())
	}
}

func TestSm3Incremental(t *testing.T) {
	h := New()
	h.Write([]byte("a"))
	h.Write([]byte("b"))
	h.Write([]byte("c"))
	got := hex.EncodeToString(h.Sum(nil))
	expected := "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0"
	if got != expected {
		t.Errorf("Incremental write: got %s, want %s", got, expected)
	}
}

func TestSm3Empty(t *testing.T) {
	got := hex.EncodeToString(Sm3Sum([]byte{}))
	expected := "1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b"
	if got != expected {
		t.Errorf("Sm3Sum(empty) = %s, want %s", got, expected)
	}
}

func TestSm3LongMessage(t *testing.T) {
	data := make([]byte, 64)
	for i := range data {
		data[i] = byte(i)
	}
	got := hex.EncodeToString(Sm3Sum(data))
	if len(got) != 64 {
		t.Errorf("Sm3Sum long message: got length %d, want 64", len(got))
	}
}
