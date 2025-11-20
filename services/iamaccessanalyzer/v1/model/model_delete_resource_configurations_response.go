package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceConfigurationsResponse Response Object
type DeleteResourceConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceConfigurationsResponse", string(data)}, " ")
}
