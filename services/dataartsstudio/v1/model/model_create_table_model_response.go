package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableModelResponse Response Object
type CreateTableModelResponse struct {
	Data           *CreateTableModelResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CreateTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableModelResponse struct{}"
	}

	return strings.Join([]string{"CreateTableModelResponse", string(data)}, " ")
}
