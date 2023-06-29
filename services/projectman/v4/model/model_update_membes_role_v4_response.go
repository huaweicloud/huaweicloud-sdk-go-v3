package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMembesRoleV4Response Response Object
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
