package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPubInfosResponse struct {
	Data           *ListPubInfosResponseModel `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListPubInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPubInfosResponse struct{}"
	}

	return strings.Join([]string{"ListPubInfosResponse", string(data)}, " ")
}
