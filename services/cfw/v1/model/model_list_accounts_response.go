package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsResponse Response Object
type ListAccountsResponse struct {
	Data           *PageDataAccountVo `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListAccountsResponse", string(data)}, " ")
}
