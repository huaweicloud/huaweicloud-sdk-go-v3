package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithIegRequest Request Object
type ListEcnWithIegRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`
}

func (o ListEcnWithIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithIegRequest struct{}"
	}

	return strings.Join([]string{"ListEcnWithIegRequest", string(data)}, " ")
}
