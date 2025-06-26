package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRegistryRequest Request Object
type ShowInstanceRegistryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步仓库ID
	RegistryId int32 `json:"registry_id"`
}

func (o ShowInstanceRegistryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRegistryRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRegistryRequest", string(data)}, " ")
}
