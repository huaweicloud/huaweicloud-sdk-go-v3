package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto query response Object
type GetUsersListdatabase struct {
	// DDM实例帐号关联的逻辑库名称。

	Name *string `json:"name,omitempty"`
	// 逻辑库的描述信息。

	Description *string `json:"description,omitempty"`
}

func (o GetUsersListdatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUsersListdatabase struct{}"
	}

	return strings.Join([]string{"GetUsersListdatabase", string(data)}, " ")
}
