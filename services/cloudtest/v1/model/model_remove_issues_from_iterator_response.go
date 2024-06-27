package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIssuesFromIteratorResponse Response Object
type RemoveIssuesFromIteratorResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o RemoveIssuesFromIteratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIssuesFromIteratorResponse struct{}"
	}

	return strings.Join([]string{"RemoveIssuesFromIteratorResponse", string(data)}, " ")
}
