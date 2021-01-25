package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowMappingRequest struct {
	Id string `json:"id"`
}

func (o KeystoneShowMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowMappingRequest", string(data)}, " ")
}
