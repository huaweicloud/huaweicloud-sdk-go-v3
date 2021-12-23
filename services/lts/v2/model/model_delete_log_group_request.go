package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteLogGroupRequest struct {
	// 日志组ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）

	LogGroupId string `json:"log_group_id"`
}

func (o DeleteLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogGroupRequest", string(data)}, " ")
}
