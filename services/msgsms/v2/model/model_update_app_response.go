package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAppResponse struct {

	// 应用KEY
	AppKey *string `json:"app_key,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用主键ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}
