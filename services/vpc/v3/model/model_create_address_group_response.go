package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAddressGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	AddressGroup   *AddressGroup `json:"address_group,omitempty" xml:"address_group"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateAddressGroupResponse", string(data)}, " ")
}
