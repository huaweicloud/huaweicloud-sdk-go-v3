package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQuotaStateRequest struct {

	// 可用区（cn-north-7a）
	AvailabilityZone string `json:"availability_zone"`

	// 资源规格（cbh.basic.50）
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ListQuotaStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaStateRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaStateRequest", string(data)}, " ")
}
