package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthorizationResponse Response Object
type ShowAuthorizationResponse struct {

	// 委托授权信息。
	Authorization  *interface{} `json:"authorization,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"ShowAuthorizationResponse", string(data)}, " ")
}
