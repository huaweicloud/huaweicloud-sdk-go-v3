package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateComplexCombineResponse Response Object
type GenerateComplexCombineResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GenerateComplexCombineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateComplexCombineResponse struct{}"
	}

	return strings.Join([]string{"GenerateComplexCombineResponse", string(data)}, " ")
}
