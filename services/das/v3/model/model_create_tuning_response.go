package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTuningResponse Response Object
type CreateTuningResponse struct {

	// 诊断信息id
	MessageId *[]string `json:"message_id,omitempty"`

	// 状态
	Status *bool `json:"status,omitempty"`

	// 诊断配额状态
	QuotaExceeded  *bool `json:"quota_exceeded,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateTuningResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTuningResponse struct{}"
	}

	return strings.Join([]string{"CreateTuningResponse", string(data)}, " ")
}
