package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceAgencyResponse Response Object
type CreateServiceAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateServiceAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateServiceAgencyResponse", string(data)}, " ")
}
