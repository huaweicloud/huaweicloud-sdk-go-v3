package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoginTypeResponse Response Object
type ShowLoginTypeResponse struct {

	// 登录类型
	LoginType      *string `json:"login_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLoginTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoginTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowLoginTypeResponse", string(data)}, " ")
}
