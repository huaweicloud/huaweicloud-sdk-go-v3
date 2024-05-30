package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableValuesResponse Response Object
type UpdateCodeTableValuesResponse struct {
	Data           *UpdateCodeTableValuesResultData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o UpdateCodeTableValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableValuesResponse struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableValuesResponse", string(data)}, " ")
}
