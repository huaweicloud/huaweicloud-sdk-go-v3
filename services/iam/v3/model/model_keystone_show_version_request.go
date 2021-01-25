package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowVersionRequest struct {
}

func (o KeystoneShowVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowVersionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowVersionRequest", string(data)}, " ")
}
