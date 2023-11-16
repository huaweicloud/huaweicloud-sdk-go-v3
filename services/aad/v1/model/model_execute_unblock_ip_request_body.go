package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteUnblockIpRequestBody struct {

	// ip地址
	Ip string `json:"ip"`

	// 用于查询IP的封堵时间
	BlockingTime *int64 `json:"blocking_time,omitempty"`
}

func (o ExecuteUnblockIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUnblockIpRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteUnblockIpRequestBody", string(data)}, " ")
}
