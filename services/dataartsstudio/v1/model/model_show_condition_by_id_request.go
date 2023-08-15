package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConditionByIdRequest Request Object
type ShowConditionByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	// 是否查询最新的
	Latest *bool `json:"latest,omitempty"`
}

func (o ShowConditionByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConditionByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowConditionByIdRequest", string(data)}, " ")
}
