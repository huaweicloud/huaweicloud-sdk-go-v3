package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResourceInstancesTag struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值列表
	Values *[]string `json:"values,omitempty"`
}

func (o TmsResourceInstancesTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceInstancesTag struct{}"
	}

	return strings.Join([]string{"TmsResourceInstancesTag", string(data)}, " ")
}
