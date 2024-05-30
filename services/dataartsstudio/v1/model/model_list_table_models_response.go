package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelsResponse Response Object
type ListTableModelsResponse struct {
	Data           *ListTableModelsResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListTableModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelsResponse struct{}"
	}

	return strings.Join([]string{"ListTableModelsResponse", string(data)}, " ")
}
