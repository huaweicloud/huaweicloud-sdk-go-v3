package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRelationsByOneCaseResponse Response Object
type DeleteRelationsByOneCaseResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o DeleteRelationsByOneCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationsByOneCaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteRelationsByOneCaseResponse", string(data)}, " ")
}
