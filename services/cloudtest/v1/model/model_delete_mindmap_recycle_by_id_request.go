package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMindmapRecycleByIdRequest Request Object
type DeleteMindmapRecycleByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 回收站ID
	Id string `json:"id"`
}

func (o DeleteMindmapRecycleByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMindmapRecycleByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteMindmapRecycleByIdRequest", string(data)}, " ")
}
