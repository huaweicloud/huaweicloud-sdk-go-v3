package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLabelValuesAomPromGetResponse struct {
	Status *string `json:"status,omitempty"`

	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLabelValuesAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelValuesAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListLabelValuesAomPromGetResponse", string(data)}, " ")
}
