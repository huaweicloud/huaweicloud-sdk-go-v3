package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Instance struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}
