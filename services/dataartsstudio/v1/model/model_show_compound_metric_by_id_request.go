package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompoundMetricByIdRequest Request Object
type ShowCompoundMetricByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	// 是否查询最新的
	Latest *bool `json:"latest,omitempty"`
}

func (o ShowCompoundMetricByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompoundMetricByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowCompoundMetricByIdRequest", string(data)}, " ")
}
