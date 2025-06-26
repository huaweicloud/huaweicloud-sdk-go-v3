package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRegistryRequest Request Object
type UpdateInstanceRegistryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步仓库ID
	RegistryId int32 `json:"registry_id"`

	Body *UpdateRegistryRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRegistryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRegistryRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRegistryRequest", string(data)}, " ")
}
