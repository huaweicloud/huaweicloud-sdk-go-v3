package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTokenResponse Response Object
type ShowTokenResponse struct {

	// 应用token
	Token          *string `json:"token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenResponse struct{}"
	}

	return strings.Join([]string{"ShowTokenResponse", string(data)}, " ")
}
