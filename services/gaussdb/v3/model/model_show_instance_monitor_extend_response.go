package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceMonitorExtendResponse struct {
	// 采集周期。  取值： 0表示实例秒级监控关闭； 1表示实例秒级监控开启，采集周期为1s； 5表示实例秒级监控开启，采集周期为5s。

	Period         *string `json:"period,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceMonitorExtendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMonitorExtendResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMonitorExtendResponse", string(data)}, " ")
}
