package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithErResponse Response Object
type ListEcnWithErResponse struct {

	// 企业连接网络与企业路由器的绑定关系列表
	EcnErRelationships *[]EcnErItem `json:"ecn_er_relationships,omitempty"`
	HttpStatusCode     int          `json:"-"`
}

func (o ListEcnWithErResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithErResponse struct{}"
	}

	return strings.Join([]string{"ListEcnWithErResponse", string(data)}, " ")
}
