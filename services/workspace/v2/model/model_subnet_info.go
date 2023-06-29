package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetInfo 子网信息。
type SubnetInfo struct {

	// 子网的网络ID。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o SubnetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetInfo struct{}"
	}

	return strings.Join([]string{"SubnetInfo", string(data)}, " ")
}
