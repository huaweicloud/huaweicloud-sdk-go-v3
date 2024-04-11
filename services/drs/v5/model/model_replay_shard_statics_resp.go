package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayShardStaticsResp 回放概览信息基于进间点的统计结果
type ReplayShardStaticsResp struct {

	// 回放时间点
	Time string `json:"time"`

	// SQL总量
	Total int64 `json:"total"`

	// SQL执行量
	Finish int64 `json:"finish"`

	// SQL异常量
	Abnormal int64 `json:"abnormal"`

	// 慢SQL数量
	Slow int64 `json:"slow"`
}

func (o ReplayShardStaticsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayShardStaticsResp struct{}"
	}

	return strings.Join([]string{"ReplayShardStaticsResp", string(data)}, " ")
}
