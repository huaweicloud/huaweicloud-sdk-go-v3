package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserByDomainResponse Response Object
type UpdateUserByDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserByDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserByDomainResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserByDomainResponse", string(data)}, " ")
}
