package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetKeepTimeRequestBody 设置回收站中的任务保留时间接口请求体
type SetKeepTimeRequestBody struct {

	// 回收站中的任务保留时间
	KeepTime int32 `json:"keep_time"`
}

func (o SetKeepTimeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKeepTimeRequestBody struct{}"
	}

	return strings.Join([]string{"SetKeepTimeRequestBody", string(data)}, " ")
}
