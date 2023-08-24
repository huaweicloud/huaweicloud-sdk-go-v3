package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSmsAppResponse Response Object
type CreateSmsAppResponse struct {

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用ID。
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSmsAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmsAppResponse struct{}"
	}

	return strings.Join([]string{"CreateSmsAppResponse", string(data)}, " ")
}
