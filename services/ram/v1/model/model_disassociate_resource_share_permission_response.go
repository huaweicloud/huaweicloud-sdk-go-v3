package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResourceSharePermissionResponse Response Object
type DisassociateResourceSharePermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateResourceSharePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResourceSharePermissionResponse struct{}"
	}

	return strings.Join([]string{"DisassociateResourceSharePermissionResponse", string(data)}, " ")
}
