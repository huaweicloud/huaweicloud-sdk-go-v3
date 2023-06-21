package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePubInfoRequest struct {

	// 服务号ID。
	PubId string `json:"pub_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *UpdatePubInfoRequestBody `json:"body,omitempty"`
}

func (o UpdatePubInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePubInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdatePubInfoRequest", string(data)}, " ")
}
