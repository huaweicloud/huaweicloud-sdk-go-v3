package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResourceRequest Request Object
type ImportResourceRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 资源名称
	Resource string `json:"resource"`

	// 业务id,比如问答模板时传入skill_id
	BusinessId *string `json:"business_id,omitempty"`

	Body *ImportResourceRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResourceRequest struct{}"
	}

	return strings.Join([]string{"ImportResourceRequest", string(data)}, " ")
}
