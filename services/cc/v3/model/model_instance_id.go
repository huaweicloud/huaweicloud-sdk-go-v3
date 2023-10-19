package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceId struct {

	// 资源ID标识符。
	InstanceId string `json:"instance_id"`
}

func (o InstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceId struct{}"
	}

	return strings.Join([]string{"InstanceId", string(data)}, " ")
}
