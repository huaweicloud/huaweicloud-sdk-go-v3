package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create response Object
type CreateUsersDetailResponses struct {
	// DDM实例帐号名称。

	Name string `json:"name"`
}

func (o CreateUsersDetailResponses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersDetailResponses struct{}"
	}

	return strings.Join([]string{"CreateUsersDetailResponses", string(data)}, " ")
}
