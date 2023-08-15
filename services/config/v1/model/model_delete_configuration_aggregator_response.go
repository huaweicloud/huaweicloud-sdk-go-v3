package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigurationAggregatorResponse Response Object
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
