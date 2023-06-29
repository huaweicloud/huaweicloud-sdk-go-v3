package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateListenerResponse Response Object
type DisassociateListenerResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateListenerResponse struct{}"
	}

	return strings.Join([]string{"DisassociateListenerResponse", string(data)}, " ")
}
