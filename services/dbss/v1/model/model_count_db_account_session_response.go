package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountDbAccountSessionResponse Response Object
type CountDbAccountSessionResponse struct {

	// 用户会话量列表
	Body           *[]ReportAccountSessionNew `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CountDbAccountSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDbAccountSessionResponse struct{}"
	}

	return strings.Join([]string{"CountDbAccountSessionResponse", string(data)}, " ")
}
