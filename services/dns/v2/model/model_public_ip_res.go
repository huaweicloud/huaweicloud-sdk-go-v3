package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicIpRes struct {

	// 弹性公网IP（EIP）所属的区域。
	Region *string `json:"region,omitempty"`

	// 弹性公网IP（EIP）的ID。
	Id *string `json:"id,omitempty"`

	// 弹性公网IP（EIP）。
	Address *string `json:"address,omitempty"`
}

func (o PublicIpRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpRes struct{}"
	}

	return strings.Join([]string{"PublicIpRes", string(data)}, " ")
}
