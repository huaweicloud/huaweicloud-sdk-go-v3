package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaInstanceExtendProductInfoResponse Response Object
type ShowKafkaInstanceExtendProductInfoResponse struct {

	// 消息引擎类型：kafka。
	Engine *string `json:"engine,omitempty"`

	// 消息引擎支持的版本。
	Versions *[]string `json:"versions,omitempty"`

	// 规格变更的产品信息。
	Products       *[]ExtendProductInfoEntity `json:"products,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowKafkaInstanceExtendProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaInstanceExtendProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaInstanceExtendProductInfoResponse", string(data)}, " ")
}
