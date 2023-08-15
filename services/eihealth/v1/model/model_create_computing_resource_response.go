package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComputingResourceResponse Response Object
type CreateComputingResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateComputingResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceResponse", string(data)}, " ")
}
