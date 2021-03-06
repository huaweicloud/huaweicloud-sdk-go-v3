package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowSecurityComplianceResponse struct {
	Config         *Config `json:"config,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneShowSecurityComplianceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceResponse", string(data)}, " ")
}
