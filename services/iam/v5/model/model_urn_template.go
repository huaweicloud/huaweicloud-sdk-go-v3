package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrnTemplate 统一资源名称模板，表示可以通过这类资源的统一资源名称对该授权项进行资源粒度的授权。
type UrnTemplate struct {
}

func (o UrnTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrnTemplate struct{}"
	}

	return strings.Join([]string{"UrnTemplate", string(data)}, " ")
}
