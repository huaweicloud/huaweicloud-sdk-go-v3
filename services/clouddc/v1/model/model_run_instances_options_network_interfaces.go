package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunInstancesOptionsNetworkInterfaces struct {

	// subnet id
	SubnetId string `json:"subnet_id"`
}

func (o RunInstancesOptionsNetworkInterfaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstancesOptionsNetworkInterfaces struct{}"
	}

	return strings.Join([]string{"RunInstancesOptionsNetworkInterfaces", string(data)}, " ")
}
