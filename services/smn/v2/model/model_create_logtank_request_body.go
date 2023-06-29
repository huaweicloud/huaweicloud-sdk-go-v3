package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankRequestBody 创建云日志信息请求体
type CreateLogtankRequestBody struct {

	// 云日志服务日志组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志服务日志流ID。
	LogStreamId string `json:"log_stream_id"`
}

func (o CreateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequestBody", string(data)}, " ")
}
