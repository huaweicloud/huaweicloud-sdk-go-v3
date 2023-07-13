package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateListenerRequestBody disassociate listener request
type DisassociateListenerRequestBody struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`
}

func (o DisassociateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateListenerRequestBody", string(data)}, " ")
}
