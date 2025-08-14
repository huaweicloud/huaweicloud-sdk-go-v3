package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTaskRequest Request Object
type SearchTaskRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 资源类型
	Resource *[]string `json:"resource,omitempty"`

	// 业务id,比如问答时传入skill_id
	BusinessId *string `json:"business_id,omitempty"`

	// 开始时间戳
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间戳
	EndTime *string `json:"end_time,omitempty"`
}

func (o SearchTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTaskRequest struct{}"
	}

	return strings.Join([]string{"SearchTaskRequest", string(data)}, " ")
}
