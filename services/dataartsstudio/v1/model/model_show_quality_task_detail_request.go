package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQualityTaskDetailRequest Request Object
type ShowQualityTaskDetailRequest struct {

	// 质量作业ID
	Id string `json:"id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowQualityTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQualityTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowQualityTaskDetailRequest", string(data)}, " ")
}
