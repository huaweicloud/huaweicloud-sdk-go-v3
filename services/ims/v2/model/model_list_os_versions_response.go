package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOsVersionsResponse struct {
	Body           *[]ListOsVersionsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListOsVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListOsVersionsResponse", string(data)}, " ")
}
