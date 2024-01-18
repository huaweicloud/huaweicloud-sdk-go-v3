package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejoinDomainResponse Response Object
type BatchRejoinDomainResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRejoinDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejoinDomainResponse struct{}"
	}

	return strings.Join([]string{"BatchRejoinDomainResponse", string(data)}, " ")
}
