package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferServiceTagResponse Response Object
type CreateInferServiceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInferServiceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferServiceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateInferServiceTagResponse", string(data)}, " ")
}
