package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantIdDef 实例所属项目ID。
type TenantIdDef struct {
}

func (o TenantIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantIdDef struct{}"
	}

	return strings.Join([]string{"TenantIdDef", string(data)}, " ")
}
