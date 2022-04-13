package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVmrResponse struct{}"
	}

	return strings.Join([]string{"DisassociateVmrResponse", string(data)}, " ")
}
