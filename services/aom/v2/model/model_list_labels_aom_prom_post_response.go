package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLabelsAomPromPostResponse struct {
	Status *string `json:"status,omitempty"`

	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLabelsAomPromPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelsAomPromPostResponse struct{}"
	}

	return strings.Join([]string{"ListLabelsAomPromPostResponse", string(data)}, " ")
}
