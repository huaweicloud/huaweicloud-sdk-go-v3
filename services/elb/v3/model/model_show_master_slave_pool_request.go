package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMasterSlavePoolRequest Request Object
type ShowMasterSlavePoolRequest struct {

	// 参数解释：后端服务器组ID。
	PoolId string `json:"pool_id"`
}

func (o ShowMasterSlavePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterSlavePoolRequest struct{}"
	}

	return strings.Join([]string{"ShowMasterSlavePoolRequest", string(data)}, " ")
}
