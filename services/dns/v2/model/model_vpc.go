package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vpc struct {

	// 关联VPC的ID。
	VpcId string `json:"vpc_id"`

	// 关联VPC所在的region。
	VpcRegion *string `json:"vpc_region,omitempty"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}
