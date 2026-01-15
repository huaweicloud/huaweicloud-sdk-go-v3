package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountDbClientSessionResponse Response Object
type CountDbClientSessionResponse struct {

	// 客户端会话统计
	Body           *[]ReportClientSessionNew `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CountDbClientSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDbClientSessionResponse struct{}"
	}

	return strings.Join([]string{"CountDbClientSessionResponse", string(data)}, " ")
}
