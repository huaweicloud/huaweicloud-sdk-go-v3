package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchLogCollectionStatusResponse Response Object
type SwitchLogCollectionStatusResponse struct {

	// 成功响应体。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchLogCollectionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchLogCollectionStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchLogCollectionStatusResponse", string(data)}, " ")
}
