package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAddressSetInfoUsingPutRequest struct {

	// 地址组id
	SetId string `json:"set_id"`

	Body *UpdateAddressSetDto `json:"body,omitempty"`
}

func (o UpdateAddressSetInfoUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetInfoUsingPutRequest struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetInfoUsingPutRequest", string(data)}, " ")
}
