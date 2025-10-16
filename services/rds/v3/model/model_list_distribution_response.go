package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDistributionResponse Response Object
type ListDistributionResponse struct {

	// 状态。normal，abnormal，creating，createfail
	Status *string `json:"status,omitempty"`

	// 分发服务器实例id。
	DistributorInstanceId *string `json:"distributor_instance_id,omitempty"`

	// 分发服务器实例名称。
	DistributorInstanceName *string `json:"distributor_instance_name,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o ListDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDistributionResponse struct{}"
	}

	return strings.Join([]string{"ListDistributionResponse", string(data)}, " ")
}
