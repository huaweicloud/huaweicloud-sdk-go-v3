package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGrantsResponse Response Object
type ListGrantsResponse struct {
	Data           *GrantDto `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGrantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGrantsResponse struct{}"
	}

	return strings.Join([]string{"ListGrantsResponse", string(data)}, " ")
}
