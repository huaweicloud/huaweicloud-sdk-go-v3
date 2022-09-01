package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupAssignment struct {

	// topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 分区列表。
	Partitions *[]int32 `json:"partitions,omitempty" xml:"partitions"`
}

func (o ShowGroupsRespGroupAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupAssignment struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupAssignment", string(data)}, " ")
}
