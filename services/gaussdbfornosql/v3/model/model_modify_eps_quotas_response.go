package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ModifyEpsQuotasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyEpsQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEpsQuotasResponse struct{}"
	}

	return strings.Join([]string{"ModifyEpsQuotasResponse", string(data)}, " ")
}
