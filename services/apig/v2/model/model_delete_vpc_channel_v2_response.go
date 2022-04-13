package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVpcChannelV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcChannelV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"DeleteVpcChannelV2Response", string(data)}, " ")
}
