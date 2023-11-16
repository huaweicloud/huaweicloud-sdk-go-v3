package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainResponse Response Object
type UpdateDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainResponse", string(data)}, " ")
}
