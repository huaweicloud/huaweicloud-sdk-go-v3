package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsRequest Request Object
type ListProductsRequest struct {

	// 消息引擎的类型。当前只支持rabbitmq。
	Engine *string `json:"engine,omitempty"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}
