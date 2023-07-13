package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSubnetRequest Request Object
type NeutronCreateSubnetRequest struct {
	Body *NeutronCreateSubnetRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSubnetRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSubnetRequest", string(data)}, " ")
}
