package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupAssignment struct {
	// topic名称。

	Topic *string `json:"topic,omitempty"`
	// 分区列表。

	Partitions *[]int32 `json:"partitions,omitempty"`
}

func (o ShowGroupsRespGroupAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupAssignment struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupAssignment", string(data)}, " ")
}
