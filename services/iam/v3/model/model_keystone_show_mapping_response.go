package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowMappingResponse struct {
	Mapping        *MappingResult `json:"mapping,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneShowMappingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowMappingResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowMappingResponse", string(data)}, " ")
}
