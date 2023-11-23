package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMasterSlavePoolResponse Response Object
type ShowMasterSlavePoolResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Pool           *MasterSlavePool `json:"pool,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowMasterSlavePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterSlavePoolResponse struct{}"
	}

	return strings.Join([]string{"ShowMasterSlavePoolResponse", string(data)}, " ")
}
