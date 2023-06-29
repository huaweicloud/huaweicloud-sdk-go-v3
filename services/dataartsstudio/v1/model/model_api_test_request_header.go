package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiTestRequestHeader struct {

	// 请求路径
	Path *string `json:"path,omitempty"`

	// 代理（固定值）
	UserAgent *string `json:"user_agent,omitempty"`

	// 请求方式（固定值）
	XApigMode *string `json:"x_apig_mode,omitempty"`

	// 识别编号（固定值）
	XAppIdentity *int32 `json:"x_app_identity,omitempty"`
}

func (o ApiTestRequestHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiTestRequestHeader struct{}"
	}

	return strings.Join([]string{"ApiTestRequestHeader", string(data)}, " ")
}
