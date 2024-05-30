package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConditionResponse Response Object
type ListConditionResponse struct {
	Data           *ListConditionResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConditionResponse struct{}"
	}

	return strings.Join([]string{"ListConditionResponse", string(data)}, " ")
}
