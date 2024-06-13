package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachmentsRequest Request Object
type ListAttachmentsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资源Uri
	ResourceUri string `json:"resource_uri"`

	// 资源类型
	ResourceType string `json:"resource_type"`
}

func (o ListAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListAttachmentsRequest", string(data)}, " ")
}
