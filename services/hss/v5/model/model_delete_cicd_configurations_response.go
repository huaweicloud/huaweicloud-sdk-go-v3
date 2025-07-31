package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCicdConfigurationsResponse Response Object
type DeleteCicdConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCicdConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCicdConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"DeleteCicdConfigurationsResponse", string(data)}, " ")
}
