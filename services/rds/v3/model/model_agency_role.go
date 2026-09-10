package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyRole struct {

	// 角色名称。
	Name *string `json:"name,omitempty"`

	// 角色描述。
	Description *string `json:"description,omitempty"`
}

func (o AgencyRole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyRole struct{}"
	}

	return strings.Join([]string{"AgencyRole", string(data)}, " ")
}
