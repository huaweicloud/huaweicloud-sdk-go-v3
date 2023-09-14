package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionUrlResponse Response Object
type ShowFunctionUrlResponse struct {

	// 函数URL鉴权方式
	AuthType *string `json:"auth_type,omitempty"`

	Cors *CorsConfig `json:"cors,omitempty"`

	// 函数URL地址
	FunctionUrl *string `json:"function_url,omitempty"`

	// 函数URL创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 函数URL上次修改时间
	LastModifiedTime *string `json:"last_modified_time,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowFunctionUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionUrlResponse", string(data)}, " ")
}
