package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMixJobRequest Request Object
type CreateMixJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用id
	AppId string `json:"app_id"`

	Body *MixJobReq `json:"body,omitempty"`
}

func (o CreateMixJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMixJobRequest struct{}"
	}

	return strings.Join([]string{"CreateMixJobRequest", string(data)}, " ")
}
