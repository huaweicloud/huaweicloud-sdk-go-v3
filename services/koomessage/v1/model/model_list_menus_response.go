package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMenusResponse Response Object
type ListMenusResponse struct {
	Data           *ListMenusResponseModel `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListMenusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMenusResponse struct{}"
	}

	return strings.Join([]string{"ListMenusResponse", string(data)}, " ")
}
