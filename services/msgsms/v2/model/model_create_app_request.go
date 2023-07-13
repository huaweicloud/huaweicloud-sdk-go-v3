package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppRequest Request Object
type CreateAppRequest struct {
	Body *SmsAppAddReq `json:"body,omitempty"`
}

func (o CreateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequest struct{}"
	}

	return strings.Join([]string{"CreateAppRequest", string(data)}, " ")
}
