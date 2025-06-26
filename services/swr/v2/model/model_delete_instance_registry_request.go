package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceRegistryRequest Request Object
type DeleteInstanceRegistryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 同步仓库ID
	RegistryId int32 `json:"registry_id"`
}

func (o DeleteInstanceRegistryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRegistryRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRegistryRequest", string(data)}, " ")
}
