package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceId 实例Id
type InstanceId struct {
}

func (o InstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceId struct{}"
	}

	return strings.Join([]string{"InstanceId", string(data)}, " ")
}
