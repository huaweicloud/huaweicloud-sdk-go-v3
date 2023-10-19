package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipsResponse Response Object
type ListEipsResponse struct {
	Data           *EipResponseData `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipsResponse struct{}"
	}

	return strings.Join([]string{"ListEipsResponse", string(data)}, " ")
}
