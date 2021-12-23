package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryAuthorizedNodeDto struct {
	// 边缘节点ID

	NodeId *string `json:"node_id,omitempty"`
	// 授权时间

	AuthorizedTime *string `json:"authorized_time,omitempty"`
}

func (o QueryAuthorizedNodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAuthorizedNodeDto struct{}"
	}

	return strings.Join([]string{"QueryAuthorizedNodeDto", string(data)}, " ")
}
