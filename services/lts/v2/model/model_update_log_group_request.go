package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogGroupRequest Request Object
type UpdateLogGroupRequest struct {

	// 日志组ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）
	LogGroupId string `json:"log_group_id"`

	Body *UpdateLogGroupParams `json:"body,omitempty"`
}

func (o UpdateLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogGroupRequest", string(data)}, " ")
}
