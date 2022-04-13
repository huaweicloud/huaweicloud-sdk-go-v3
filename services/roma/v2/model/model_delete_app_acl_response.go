package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAppAclResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppAclResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppAclResponse", string(data)}, " ")
}
