package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateSqlRequest struct {
	Body *ValidateSqlRequestBody `json:"body,omitempty"`
}

func (o ValidateSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSqlRequest struct{}"
	}

	return strings.Join([]string{"ValidateSqlRequest", string(data)}, " ")
}
