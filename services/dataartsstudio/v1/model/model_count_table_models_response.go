package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTableModelsResponse Response Object
type CountTableModelsResponse struct {
	Data           *CountTableModelsResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CountTableModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTableModelsResponse struct{}"
	}

	return strings.Join([]string{"CountTableModelsResponse", string(data)}, " ")
}
