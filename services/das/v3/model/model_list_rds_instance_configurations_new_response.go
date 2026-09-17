package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRdsInstanceConfigurationsNewResponse Response Object
type ListRdsInstanceConfigurationsNewResponse struct {
	Body           *[]ConfigurationParameterDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListRdsInstanceConfigurationsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRdsInstanceConfigurationsNewResponse struct{}"
	}

	return strings.Join([]string{"ListRdsInstanceConfigurationsNewResponse", string(data)}, " ")
}
