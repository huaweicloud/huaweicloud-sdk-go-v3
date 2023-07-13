package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateSubnetRequest Request Object
type NeutronUpdateSubnetRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`

	Body *NeutronUpdateSubnetRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSubnetRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSubnetRequest", string(data)}, " ")
}
