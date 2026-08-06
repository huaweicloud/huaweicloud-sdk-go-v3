package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelShareNewResponse Response Object
type CancelShareNewResponse struct {

	// 操作是否成功
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CancelShareNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareNewResponse struct{}"
	}

	return strings.Join([]string{"CancelShareNewResponse", string(data)}, " ")
}
