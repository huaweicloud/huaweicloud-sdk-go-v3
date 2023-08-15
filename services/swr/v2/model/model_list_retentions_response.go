package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetentionsResponse Response Object
type ListRetentionsResponse struct {
	Body           *[]Retention `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRetentionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetentionsResponse struct{}"
	}

	return strings.Join([]string{"ListRetentionsResponse", string(data)}, " ")
}
