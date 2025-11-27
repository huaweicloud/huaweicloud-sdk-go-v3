package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceGroupRequestBody struct {

	// 实例组ID
	GroupId int32 `json:"group_id"`

	// 实例组名称
	GroupName *string `json:"group_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateInstanceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceGroupRequestBody", string(data)}, " ")
}
