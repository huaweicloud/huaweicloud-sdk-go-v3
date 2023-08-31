package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnAccessPointRequest Request Object
type DeleteEcnAccessPointRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 企业连接网络接入点ID
	AccessPointId string `json:"access_point_id"`
}

func (o DeleteEcnAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnAccessPointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEcnAccessPointRequest", string(data)}, " ")
}
