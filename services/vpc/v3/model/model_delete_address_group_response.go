package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressGroupResponse Response Object
type DeleteAddressGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressGroupResponse", string(data)}, " ")
}
