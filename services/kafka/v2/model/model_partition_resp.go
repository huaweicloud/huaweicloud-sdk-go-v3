package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionResp 分区响应信息
type PartitionResp struct {

	// 分区
	Partition *int32 `json:"partition,omitempty"`

	// 返回结果
	Result *string `json:"result,omitempty"`

	// 返回错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o PartitionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionResp struct{}"
	}

	return strings.Join([]string{"PartitionResp", string(data)}, " ")
}
