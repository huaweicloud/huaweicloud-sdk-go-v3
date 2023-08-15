package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankRequestBody 更新云日志信息请求体
type UpdateLogtankRequestBody struct {

	// 云日志服务日志组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志服务日志流ID。
	LogStreamId string `json:"log_stream_id"`
}

func (o UpdateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLogtankRequestBody", string(data)}, " ")
}
