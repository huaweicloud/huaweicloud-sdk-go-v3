package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapCreatorNameRequest Request Object
type ShowMindmapCreatorNameRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ShowMindmapCreatorNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapCreatorNameRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapCreatorNameRequest", string(data)}, " ")
}
