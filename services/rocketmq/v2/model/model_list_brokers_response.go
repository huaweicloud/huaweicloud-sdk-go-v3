package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBrokersResponse struct {

	// 代理列表。
	Brokers        *[]ListBrokersRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListBrokersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersResponse struct{}"
	}

	return strings.Join([]string{"ListBrokersResponse", string(data)}, " ")
}
