package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainNewResponse Response Object
type UpdateDomainNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainNewResponse", string(data)}, " ")
}
