package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceExtendProductInfoResponse Response Object
type ShowInstanceExtendProductInfoResponse struct {

	// 消息引擎类型。
	Engine *string `json:"engine,omitempty"`

	// 消息引擎支持的版本
	Versions *[]string `json:"versions,omitempty"`

	// 规格变更的产品信息。
	Products       *[]RabbitMqExtendProductInfoEntity `json:"products,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowInstanceExtendProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoResponse", string(data)}, " ")
}
