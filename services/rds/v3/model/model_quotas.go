package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quotas struct {
	// 资源列表对象。

	Resources *[]Resources `json:"resources,omitempty"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
