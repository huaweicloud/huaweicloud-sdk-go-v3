package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationsResponse Response Object
type ListRelationsResponse struct {
	Data           *ListRelationsResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListRelationsResponse", string(data)}, " ")
}
