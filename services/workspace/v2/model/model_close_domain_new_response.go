package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseDomainNewResponse Response Object
type CloseDomainNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseDomainNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseDomainNewResponse struct{}"
	}

	return strings.Join([]string{"CloseDomainNewResponse", string(data)}, " ")
}
