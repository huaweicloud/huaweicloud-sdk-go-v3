package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDirectConnectRequestBody 物理专线更新对象参数
type UpdateDirectConnectRequestBody struct {
	DirectConnect *UpdateDirectConnect `json:"direct_connect,omitempty"`
}

func (o UpdateDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectRequestBody", string(data)}, " ")
}
