package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithErRequest Request Object
type ListEcnWithErRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`
}

func (o ListEcnWithErRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithErRequest struct{}"
	}

	return strings.Join([]string{"ListEcnWithErRequest", string(data)}, " ")
}
