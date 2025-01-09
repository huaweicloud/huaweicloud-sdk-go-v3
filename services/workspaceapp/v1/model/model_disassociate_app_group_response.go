package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppGroupResponse Response Object
type DisassociateAppGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppGroupResponse struct{}"
	}

	return strings.Join([]string{"DisassociateAppGroupResponse", string(data)}, " ")
}
