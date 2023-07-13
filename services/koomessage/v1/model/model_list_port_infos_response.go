package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortInfosResponse Response Object
type ListPortInfosResponse struct {
	Data           *ListPortInfosResponseModel `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListPortInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortInfosResponse struct{}"
	}

	return strings.Join([]string{"ListPortInfosResponse", string(data)}, " ")
}
