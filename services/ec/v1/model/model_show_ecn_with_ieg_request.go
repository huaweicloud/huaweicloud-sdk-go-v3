package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEcnWithIegRequest Request Object
type ShowEcnWithIegRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 企业连接网络绑定关系ID
	RelationId string `json:"relation_id"`
}

func (o ShowEcnWithIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcnWithIegRequest struct{}"
	}

	return strings.Join([]string{"ShowEcnWithIegRequest", string(data)}, " ")
}
