package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePartitionReq struct {
	// 期望调整分区后的数量，必须大于当前分区数量，小于等于100。

	Partition *int32 `json:"partition,omitempty"`
}

func (o CreatePartitionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionReq struct{}"
	}

	return strings.Join([]string{"CreatePartitionReq", string(data)}, " ")
}
