package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlSwitchStatusResponse Response Object
type ShowSqlSwitchStatusResponse struct {

	// 开关状态。取值： Enabled：已开启， Disabled：已关闭， Switching：开关切换中
	Status *string `json:"status,omitempty"`

	// SQL数据保存天数。
	RetentionDays  *int64 `json:"retention_days,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlSwitchStatusResponse", string(data)}, " ")
}
