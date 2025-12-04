package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketsResponse Response Object
type ListObsBucketsResponse struct {
	Body           *[]ListObsBucketsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListObsBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketsResponse", string(data)}, " ")
}
