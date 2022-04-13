package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSubnetResponse struct {
	Subnet         *UpdateSubnetResponseObject `json:"subnet,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponse", string(data)}, " ")
}
