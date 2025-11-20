package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceConfigurationsResponse Response Object
type CreateResourceConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateResourceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceConfigurationsResponse", string(data)}, " ")
}
