package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplaceDefinerInfo 设置replaceDefiner信息
type ReplaceDefinerInfo struct {

	// 任务id
	JobId string `json:"job_id"`

	// 是否使用目标库的用户替换掉definer
	ReplaceDefiner bool `json:"replace_definer"`
}

func (o ReplaceDefinerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaceDefinerInfo struct{}"
	}

	return strings.Join([]string{"ReplaceDefinerInfo", string(data)}, " ")
}
