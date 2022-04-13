package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateVmrResponse struct{}"
	}

	return strings.Join([]string{"AssociateVmrResponse", string(data)}, " ")
}
