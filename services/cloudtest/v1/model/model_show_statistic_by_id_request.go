package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticByIdRequest Request Object
type ShowStatisticByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 脑图ID
	MindmapId string `json:"mindmap_id"`
}

func (o ShowStatisticByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticByIdRequest", string(data)}, " ")
}
