package model

import (
	"encoding/json"

	"strings"
)

// 设置replaceDefiner信息
type ReplaceDefinerInfo struct {
	// 任务id

	JobId string `json:"job_id"`
	// 是否使用目标库的用户替换掉definer

	ReplaceDefiner bool `json:"replace_definer"`
}

func (o ReplaceDefinerInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReplaceDefinerInfo struct{}"
	}

	return strings.Join([]string{"ReplaceDefinerInfo", string(data)}, " ")
}
