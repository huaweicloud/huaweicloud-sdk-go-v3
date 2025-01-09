package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableIpRequest Request Object
type ShowAvailableIpRequest struct {

	// 子网id。
	SubnetId string `json:"subnet_id"`
}

func (o ShowAvailableIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableIpRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableIpRequest", string(data)}, " ")
}
