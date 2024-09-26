package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainIdDef 实例所属账号ID。
type DomainIdDef struct {
}

func (o DomainIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainIdDef struct{}"
	}

	return strings.Join([]string{"DomainIdDef", string(data)}, " ")
}
