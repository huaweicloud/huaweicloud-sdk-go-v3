package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEcnRequest Request Object
type UpdateEcnRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	Body *UpdateEcnRequestBody `json:"body,omitempty"`
}

func (o UpdateEcnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnRequest struct{}"
	}

	return strings.Join([]string{"UpdateEcnRequest", string(data)}, " ")
}
