package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEcnAccessPointRequest Request Object
type UpdateEcnAccessPointRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 企业连接网络接入点ID
	AccessPointId string `json:"access_point_id"`

	Body *UpdateAccessPointRequestBody `json:"body,omitempty"`
}

func (o UpdateEcnAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnAccessPointRequest struct{}"
	}

	return strings.Join([]string{"UpdateEcnAccessPointRequest", string(data)}, " ")
}
