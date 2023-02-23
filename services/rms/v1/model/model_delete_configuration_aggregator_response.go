package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteConfigurationAggregatorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConfigurationAggregatorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationAggregatorResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationAggregatorResponse", string(data)}, " ")
}
