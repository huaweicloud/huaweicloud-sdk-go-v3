package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddRelationsByOneCaseResponse Response Object
type BatchAddRelationsByOneCaseResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchAddRelationsByOneCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddRelationsByOneCaseResponse struct{}"
	}

	return strings.Join([]string{"BatchAddRelationsByOneCaseResponse", string(data)}, " ")
}
