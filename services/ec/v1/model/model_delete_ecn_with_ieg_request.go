package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnWithIegRequest Request Object
type DeleteEcnWithIegRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 企业连接网络绑定关系ID
	RelationId string `json:"relation_id"`
}

func (o DeleteEcnWithIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnWithIegRequest struct{}"
	}

	return strings.Join([]string{"DeleteEcnWithIegRequest", string(data)}, " ")
}
