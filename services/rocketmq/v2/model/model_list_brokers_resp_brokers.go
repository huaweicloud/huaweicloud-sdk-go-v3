package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListBrokersRespBrokers struct {

	// 全部代理ID。
	Ids *[]float32 `json:"ids,omitempty"`

	// 节点名称。
	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ListBrokersRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRespBrokers struct{}"
	}

	return strings.Join([]string{"ListBrokersRespBrokers", string(data)}, " ")
}
