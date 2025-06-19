package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetKeepTimeResult 返回结果
type SetKeepTimeResult struct {

	// 回收站中的任务保留时间
	KeepTime *int32 `json:"keep_time,omitempty"`
}

func (o SetKeepTimeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKeepTimeResult struct{}"
	}

	return strings.Join([]string{"SetKeepTimeResult", string(data)}, " ")
}
