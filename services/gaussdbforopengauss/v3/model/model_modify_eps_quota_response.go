package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ModifyEpsQuotaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyEpsQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEpsQuotaResponse struct{}"
	}

	return strings.Join([]string{"ModifyEpsQuotaResponse", string(data)}, " ")
}
