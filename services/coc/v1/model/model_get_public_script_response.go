package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPublicScriptResponse Response Object
type GetPublicScriptResponse struct {
	Data           *PublicScriptDetailModel `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o GetPublicScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPublicScriptResponse struct{}"
	}

	return strings.Join([]string{"GetPublicScriptResponse", string(data)}, " ")
}
