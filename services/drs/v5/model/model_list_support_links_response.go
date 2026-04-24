package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportLinksResponse Response Object
type ListSupportLinksResponse struct {
	Body           *[]SupportLinksResp `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSupportLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportLinksResponse struct{}"
	}

	return strings.Join([]string{"ListSupportLinksResponse", string(data)}, " ")
}
