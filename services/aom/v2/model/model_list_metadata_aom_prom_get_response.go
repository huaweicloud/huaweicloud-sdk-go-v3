package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMetadataAomPromGetResponse struct {
	Status *string `json:"status,omitempty"`

	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListMetadataAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadataAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListMetadataAomPromGetResponse", string(data)}, " ")
}
