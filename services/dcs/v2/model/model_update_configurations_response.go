package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigurationsResponse Response Object
type UpdateConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationsResponse", string(data)}, " ")
}
