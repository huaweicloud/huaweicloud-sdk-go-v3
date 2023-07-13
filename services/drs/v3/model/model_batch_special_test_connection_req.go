package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSpecialTestConnectionReq 集群批量测试连接任务请求体
type BatchSpecialTestConnectionReq struct {

	// 集群批量测试连接任务请求列表
	Jobs []BatchJobActionReq `json:"jobs"`
}

func (o BatchSpecialTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSpecialTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchSpecialTestConnectionReq", string(data)}, " ")
}
