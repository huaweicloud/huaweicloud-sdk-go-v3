package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicIpReq struct {

	// 弹性公网IP（EIP）所属的区域。
	Region *string `json:"region,omitempty"`

	// 弹性公网IP（EIP）的ID。
	Id *string `json:"id,omitempty"`
}

func (o PublicIpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpReq struct{}"
	}

	return strings.Join([]string{"PublicIpReq", string(data)}, " ")
}
