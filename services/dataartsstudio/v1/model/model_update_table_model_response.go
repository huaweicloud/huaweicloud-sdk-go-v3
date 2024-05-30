package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableModelResponse Response Object
type UpdateTableModelResponse struct {
	Data           *CreateTableModelResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateTableModelResponse", string(data)}, " ")
}
