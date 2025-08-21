package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImpersonationTokensResponse Response Object
type ListImpersonationTokensResponse struct {
	Body           *[]ImpersonationToken `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListImpersonationTokensResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImpersonationTokensResponse struct{}"
	}

	return strings.Join([]string{"ListImpersonationTokensResponse", string(data)}, " ")
}
