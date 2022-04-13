package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
