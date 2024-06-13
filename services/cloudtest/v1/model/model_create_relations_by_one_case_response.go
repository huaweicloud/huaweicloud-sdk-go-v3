package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRelationsByOneCaseResponse Response Object
type CreateRelationsByOneCaseResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateRelationsByOneCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRelationsByOneCaseResponse struct{}"
	}

	return strings.Join([]string{"CreateRelationsByOneCaseResponse", string(data)}, " ")
}
