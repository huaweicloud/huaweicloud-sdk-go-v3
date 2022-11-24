package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProtectableResponse struct {
	Instance       *ProtectablesResp `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProtectableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectableResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectableResponse", string(data)}, " ")
}
