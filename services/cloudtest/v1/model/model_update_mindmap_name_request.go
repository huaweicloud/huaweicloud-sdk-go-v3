package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMindmapNameRequest Request Object
type UpdateMindmapNameRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 脑图ID
	Id string `json:"id"`

	// 脑图名称
	Name string `json:"name"`
}

func (o UpdateMindmapNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMindmapNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateMindmapNameRequest", string(data)}, " ")
}
