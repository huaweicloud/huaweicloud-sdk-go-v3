package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Dependent struct {
	// 部署方式。

	Name *string `json:"name,omitempty"`
	// 依赖云资源信息

	DependentServices *[]ResouceInfo `json:"dependent_services,omitempty"`
}

func (o Dependent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dependent struct{}"
	}

	return strings.Join([]string{"Dependent", string(data)}, " ")
}
