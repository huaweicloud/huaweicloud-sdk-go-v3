package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezePubRequest Request Object
type FreezePubRequest struct {

	// 服务号ID。
	PubId string `json:"pub_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *FreezePubRequestBody `json:"body,omitempty"`
}

func (o FreezePubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubRequest struct{}"
	}

	return strings.Join([]string{"FreezePubRequest", string(data)}, " ")
}
