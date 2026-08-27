package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtocolConfigsResponse Response Object
type ListProtocolConfigsResponse struct {

	// 泛协议配置列表
	ProtocolConfigs *[]ProtocolConfigBase `json:"protocol_configs,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListProtocolConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtocolConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListProtocolConfigsResponse", string(data)}, " ")
}
