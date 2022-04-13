package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMembesRoleV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMembesRoleV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMembesRoleV4Response struct{}"
	}

	return strings.Join([]string{"UpdateMembesRoleV4Response", string(data)}, " ")
}
