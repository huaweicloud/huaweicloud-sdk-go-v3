package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateEipResponse Response Object
type DisassociateEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEipResponse struct{}"
	}

	return strings.Join([]string{"DisassociateEipResponse", string(data)}, " ")
}
