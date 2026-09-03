package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelConnectionProcessResponse Response Object
type CancelConnectionProcessResponse struct {

	// 查杀的会话的数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CancelConnectionProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelConnectionProcessResponse struct{}"
	}

	return strings.Join([]string{"CancelConnectionProcessResponse", string(data)}, " ")
}
