package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddInstanceToGroupRequestBody struct {

	// 实例组ID
	GroupId int32 `json:"group_id"`

	// 实例ID列表
	InstanceIds []string `json:"instance_ids"`
}

func (o AddInstanceToGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceToGroupRequestBody struct{}"
	}

	return strings.Join([]string{"AddInstanceToGroupRequestBody", string(data)}, " ")
}
