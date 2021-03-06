package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowIdentityProviderRequest struct {
	// 待查询的身份提供商ID。

	Id string `json:"id"`
}

func (o KeystoneShowIdentityProviderRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowIdentityProviderRequest", string(data)}, " ")
}
