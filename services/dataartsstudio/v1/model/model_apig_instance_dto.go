package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApigInstanceDto struct {

	// 网关实例编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 网关实例名称
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o ApigInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigInstanceDto struct{}"
	}

	return strings.Join([]string{"ApigInstanceDto", string(data)}, " ")
}
