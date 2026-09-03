package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableQuotaResponse Response Object
type EnableQuotaResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o EnableQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableQuotaResponse struct{}"
	}

	return strings.Join([]string{"EnableQuotaResponse", string(data)}, " ")
}
