package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyCustomizedFieldsResponse Response Object
type ModifyCustomizedFieldsResponse struct {
	Data           *ModifyCustomizedFieldsResultData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ModifyCustomizedFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyCustomizedFieldsResponse struct{}"
	}

	return strings.Join([]string{"ModifyCustomizedFieldsResponse", string(data)}, " ")
}
