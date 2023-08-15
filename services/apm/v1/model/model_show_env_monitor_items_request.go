package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvMonitorItemsRequest Request Object
type ShowEnvMonitorItemsRequest struct {

	// 环境id。
	EnvId int64 `json:"env_id"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowEnvMonitorItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvMonitorItemsRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvMonitorItemsRequest", string(data)}, " ")
}
