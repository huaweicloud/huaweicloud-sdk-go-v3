package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListL7policiesResponse struct {
	// 转发策略对象的列表

	L7policies     *[]L7policyResp `json:"l7policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListL7policiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7policiesResponse struct{}"
	}

	return strings.Join([]string{"ListL7policiesResponse", string(data)}, " ")
}
