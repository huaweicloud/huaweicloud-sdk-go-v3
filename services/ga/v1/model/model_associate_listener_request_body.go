package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateListenerRequestBody associate listener request
type AssociateListenerRequestBody struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`

	Type *ListenerAccessControlType `json:"type"`
}

func (o AssociateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateListenerRequestBody", string(data)}, " ")
}
