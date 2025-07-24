package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceIpRequest Request Object
type ModifyInstanceIpRequest struct {

	// imetal instance id
	Id string `json:"id"`

	Body *ModifyInstanceIpRequestBody `json:"body,omitempty"`
}

func (o ModifyInstanceIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceIpRequest struct{}"
	}

	return strings.Join([]string{"ModifyInstanceIpRequest", string(data)}, " ")
}
