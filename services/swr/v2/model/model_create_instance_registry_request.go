package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRegistryRequest Request Object
type CreateInstanceRegistryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateRegistryRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceRegistryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRegistryRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRegistryRequest", string(data)}, " ")
}
