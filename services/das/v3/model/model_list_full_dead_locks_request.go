package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullDeadLocksRequest Request Object
type ListFullDeadLocksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间戳（ms）
	StartAt int64 `json:"start_at"`

	// 结束时间戳（ms）
	EndAt int64 `json:"end_at"`

	// 当前页
	PageNum int32 `json:"page_num"`

	// 分页大小
	PageSize int32 `json:"page_size"`
}

func (o ListFullDeadLocksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullDeadLocksRequest struct{}"
	}

	return strings.Join([]string{"ListFullDeadLocksRequest", string(data)}, " ")
}
