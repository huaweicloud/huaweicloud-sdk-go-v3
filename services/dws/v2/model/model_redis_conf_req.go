package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConfReq 重分布配置请求信息
type RedisConfReq struct {

	// 重分布模式。offLine和onLine，默认offLine。
	RedisMode string `json:"redis_mode"`

	// 并行作业数量。可配置并发数在1-200之间，默认值：4。
	ParallelJobs int32 `json:"parallel_jobs"`
}

func (o RedisConfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConfReq struct{}"
	}

	return strings.Join([]string{"RedisConfReq", string(data)}, " ")
}
