package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferServiceTagResponse Response Object
type DeleteInferServiceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInferServiceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferServiceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteInferServiceTagResponse", string(data)}, " ")
}
