package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApproverResponse Response Object
type CreateApproverResponse struct {
	Data           *CreateApproverResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CreateApproverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApproverResponse struct{}"
	}

	return strings.Join([]string{"CreateApproverResponse", string(data)}, " ")
}
