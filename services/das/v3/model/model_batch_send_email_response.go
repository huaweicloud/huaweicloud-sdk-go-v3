package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendEmailResponse Response Object
type BatchSendEmailResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchSendEmailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendEmailResponse struct{}"
	}

	return strings.Join([]string{"BatchSendEmailResponse", string(data)}, " ")
}
