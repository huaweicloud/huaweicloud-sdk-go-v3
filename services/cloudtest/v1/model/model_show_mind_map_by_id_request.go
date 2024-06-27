package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindMapByIdRequest Request Object
type ShowMindMapByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 脑图ID
	Id string `json:"id"`
}

func (o ShowMindMapByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindMapByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowMindMapByIdRequest", string(data)}, " ")
}
