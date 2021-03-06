package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowProtocolResponse struct {
	Protocol       *ProtocolResult `json:"protocol,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneShowProtocolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowProtocolResponse", string(data)}, " ")
}
