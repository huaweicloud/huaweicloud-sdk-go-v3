package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogReplayDatabaseReq 库回放请求体
type LogReplayDatabaseReq struct {

	// 需要恢复的库名列表
	Databases []RestoreInfo `json:"databases"`
}

func (o LogReplayDatabaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogReplayDatabaseReq struct{}"
	}

	return strings.Join([]string{"LogReplayDatabaseReq", string(data)}, " ")
}
