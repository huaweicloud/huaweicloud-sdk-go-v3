package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapDomainName 所属租户名称
type IsapDomainName struct {
}

func (o IsapDomainName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapDomainName struct{}"
	}

	return strings.Join([]string{"IsapDomainName", string(data)}, " ")
}
