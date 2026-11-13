package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQuotaNewResponse Response Object
type ChangeQuotaNewResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ChangeQuotaNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQuotaNewResponse struct{}"
	}

	return strings.Join([]string{"ChangeQuotaNewResponse", string(data)}, " ")
}
