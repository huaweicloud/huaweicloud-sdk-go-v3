package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecificationsResponse Response Object
type ListSpecificationsResponse struct {
	Body           *[]CbsGetSpecInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSpecificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecificationsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecificationsResponse", string(data)}, " ")
}
