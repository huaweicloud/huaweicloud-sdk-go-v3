package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicIpRequest Request Object
type ShowPublicIpRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 分页参数：起始值 [大于等于0] 。默认值是0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条 [大于0且小于等于128]。默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicIpRequest", string(data)}, " ")
}
