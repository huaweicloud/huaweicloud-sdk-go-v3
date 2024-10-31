package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapRecycleByIdRequest Request Object
type ShowMindmapRecycleByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 回收站ID
	Id string `json:"id"`
}

func (o ShowMindmapRecycleByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapRecycleByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapRecycleByIdRequest", string(data)}, " ")
}
