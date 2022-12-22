package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重分布请求
type RedistributionReq struct {

	// 重分布模式
	RedisMode string `json:"redis_mode"`

	// 重分布并发数
	ParallelJobs int32 `json:"parallel_jobs"`
}

func (o RedistributionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedistributionReq struct{}"
	}

	return strings.Join([]string{"RedistributionReq", string(data)}, " ")
}
