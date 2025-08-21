package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableMultiAccountResponse Response Object
type EnableMultiAccountResponse struct {
	Data           *EnableMultiAccountRespData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o EnableMultiAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableMultiAccountResponse struct{}"
	}

	return strings.Join([]string{"EnableMultiAccountResponse", string(data)}, " ")
}
