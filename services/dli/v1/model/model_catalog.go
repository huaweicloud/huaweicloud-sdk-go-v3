package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Catalog struct {

	// DLI侧catalog映射名称。
	Name *string `json:"name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 属性中包含type和externalCatalog
	Parameters map[string]string `json:"parameters,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o Catalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Catalog struct{}"
	}

	return strings.Join([]string{"Catalog", string(data)}, " ")
}
