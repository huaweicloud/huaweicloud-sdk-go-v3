package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAutoEnlargePolicyRequest struct {

	// 实例Id，可以调用5.3.3 查询实例列表和详情接口获取。如果未申请实例，可以调用5.3.1 创建实例接口创建。
	InstanceId string `json:"instance_id"`
}

func (o ShowAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyRequest", string(data)}, " ")
}
