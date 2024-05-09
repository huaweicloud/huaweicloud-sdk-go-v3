package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcEgressResponse Response Object
type DeleteVpcEgressResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcEgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcEgressResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcEgressResponse", string(data)}, " ")
}
