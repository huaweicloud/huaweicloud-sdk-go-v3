package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSqlRequest Request Object
type CheckSqlRequest struct {
	Body *CheckSqlRequestBody `json:"body,omitempty"`
}

func (o CheckSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSqlRequest struct{}"
	}

	return strings.Join([]string{"CheckSqlRequest", string(data)}, " ")
}
