package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnAccessPointByEcnIdRequest Request Object
type ListEcnAccessPointByEcnIdRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`
}

func (o ListEcnAccessPointByEcnIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnAccessPointByEcnIdRequest struct{}"
	}

	return strings.Join([]string{"ListEcnAccessPointByEcnIdRequest", string(data)}, " ")
}
