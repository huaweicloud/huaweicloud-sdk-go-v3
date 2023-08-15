package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezePubRequest Request Object
type UnfreezePubRequest struct {

	// 服务号ID。
	PubId string `json:"pub_id"`

	Body *UnfreezePubRequestBody `json:"body,omitempty"`
}

func (o UnfreezePubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezePubRequest struct{}"
	}

	return strings.Join([]string{"UnfreezePubRequest", string(data)}, " ")
}
