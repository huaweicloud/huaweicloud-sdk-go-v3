package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagListErrorItem struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 错误码
	ErrorCode string `json:"error_code"`

	// 错误信息
	ErrorMsg string `json:"error_msg"`
}

func (o TagListErrorItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagListErrorItem struct{}"
	}

	return strings.Join([]string{"TagListErrorItem", string(data)}, " ")
}
