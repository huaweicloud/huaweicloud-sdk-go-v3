package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstantQueryAomPromPostResponse struct {
	Status *string `json:"status,omitempty"`

	Data           *Data `json:"data,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListInstantQueryAomPromPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstantQueryAomPromPostResponse struct{}"
	}

	return strings.Join([]string{"ListInstantQueryAomPromPostResponse", string(data)}, " ")
}
