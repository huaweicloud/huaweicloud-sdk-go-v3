package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAndDeletePrivilegeReq struct {

	// 执行的操作(create|delete)
	Operation string `json:"operation" xml:"operation"`

	// 权限标识
	Privilege *string `json:"privilege,omitempty" xml:"privilege"`
}

func (o CreateAndDeletePrivilegeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndDeletePrivilegeReq struct{}"
	}

	return strings.Join([]string{"CreateAndDeletePrivilegeReq", string(data)}, " ")
}
