package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQuotaNewRequest Request Object
type ChangeQuotaNewRequest struct {
	Body *ChangeQuotaNewRequestBody `json:"body,omitempty"`
}

func (o ChangeQuotaNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQuotaNewRequest struct{}"
	}

	return strings.Join([]string{"ChangeQuotaNewRequest", string(data)}, " ")
}
