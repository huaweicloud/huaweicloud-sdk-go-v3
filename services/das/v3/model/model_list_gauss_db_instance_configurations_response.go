package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussDbInstanceConfigurationsResponse Response Object
type ListGaussDbInstanceConfigurationsResponse struct {
	Body           *[]ParameterValuesInfo `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListGaussDbInstanceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussDbInstanceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListGaussDbInstanceConfigurationsResponse", string(data)}, " ")
}
