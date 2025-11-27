package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeFederationResponse Response Object
type UpgradeFederationResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpgradeFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeFederationResponse struct{}"
	}

	return strings.Join([]string{"UpgradeFederationResponse", string(data)}, " ")
}
