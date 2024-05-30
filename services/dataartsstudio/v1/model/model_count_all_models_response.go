package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllModelsResponse Response Object
type CountAllModelsResponse struct {
	Data           *CountAllModelsResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CountAllModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllModelsResponse struct{}"
	}

	return strings.Join([]string{"CountAllModelsResponse", string(data)}, " ")
}
