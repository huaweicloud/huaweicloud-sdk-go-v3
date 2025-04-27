package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouterForRule struct {

	// Router(VPC)所属VPC的ID。
	RouterId string `json:"router_id"`
}

func (o RouterForRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterForRule struct{}"
	}

	return strings.Join([]string{"RouterForRule", string(data)}, " ")
}
