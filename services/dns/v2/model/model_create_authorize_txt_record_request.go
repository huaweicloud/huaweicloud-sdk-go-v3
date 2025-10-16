package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizeTxtRecordRequest Request Object
type CreateAuthorizeTxtRecordRequest struct {
	Body *CreateAuthorizeTxtRecordRequestBody `json:"body,omitempty"`
}

func (o CreateAuthorizeTxtRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizeTxtRecordRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorizeTxtRecordRequest", string(data)}, " ")
}
