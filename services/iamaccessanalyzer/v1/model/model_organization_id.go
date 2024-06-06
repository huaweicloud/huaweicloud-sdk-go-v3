package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationId 组织ID。
type OrganizationId struct {
}

func (o OrganizationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationId struct{}"
	}

	return strings.Join([]string{"OrganizationId", string(data)}, " ")
}
