package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource 资源。
type Resource struct {

	// 云服务资源类型名称。
	TypeName string `json:"type_name"`

	// 统一资源名称模板，表示可以通过这类资源的统一资源名称对该授权项进行资源粒度的授权。
	UrnTemplate string `json:"urn_template"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
