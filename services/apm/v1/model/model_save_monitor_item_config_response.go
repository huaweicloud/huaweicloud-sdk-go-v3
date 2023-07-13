package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveMonitorItemConfigResponse Response Object
type SaveMonitorItemConfigResponse struct {

	// 保存监控项返回状态。
	Flag           *string `json:"flag,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveMonitorItemConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveMonitorItemConfigResponse struct{}"
	}

	return strings.Join([]string{"SaveMonitorItemConfigResponse", string(data)}, " ")
}
