package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithIegResponse Response Object
type ListEcnWithIegResponse struct {

	// 企业连接网络与智能企业网关的绑定关系列表
	EcnIegRelationships *[]EcnIegItem `json:"ecn_ieg_relationships,omitempty"`
	HttpStatusCode      int           `json:"-"`
}

func (o ListEcnWithIegResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithIegResponse struct{}"
	}

	return strings.Join([]string{"ListEcnWithIegResponse", string(data)}, " ")
}
