package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLinkResponse struct {

	// 连接名称
	Name *string `json:"name,omitempty"`

	// 校验结构：如果创建连接失败，返回失败原因，请参见validation-result参数说明。如果创建成功，返回空列表。
	ValidationResult *[]ValidationResult `json:"validation-result,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o CreateLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkResponse struct{}"
	}

	return strings.Join([]string{"CreateLinkResponse", string(data)}, " ")
}
