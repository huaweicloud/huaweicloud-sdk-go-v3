package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainByDuplicateResponse Response Object
type CreateDomainByDuplicateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDomainByDuplicateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainByDuplicateResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainByDuplicateResponse", string(data)}, " ")
}
