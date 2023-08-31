package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEcnAccessPointRequest Request Object
type CreateEcnAccessPointRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	Body *CreateAccessPointRequestBody `json:"body,omitempty"`
}

func (o CreateEcnAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEcnAccessPointRequest struct{}"
	}

	return strings.Join([]string{"CreateEcnAccessPointRequest", string(data)}, " ")
}
