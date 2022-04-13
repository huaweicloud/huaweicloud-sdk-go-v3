package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMixJobRequest struct {
	// 使用AK/SK方式认证时必选，携带的鉴权信息。

	Authorization *string `json:"Authorization,omitempty"`
	// 使用AK/SK方式认证时必选，请求的发生时间。

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
	// 使用AK/SK方式认证时必选，携带项目ID信息。

	XProjectId *string `json:"X-Project-Id,omitempty"`
	// 应用id

	AppId string `json:"app_id"`
	// 任务编号

	JobId string `json:"job_id"`

	Body *UpdateMixJobReq `json:"body,omitempty"`
}

func (o UpdateMixJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMixJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateMixJobRequest", string(data)}, " ")
}
