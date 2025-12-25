package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLayoutFieldAllResponse Response Object
type ListLayoutFieldAllResponse struct {
	Body *[]LayoutFieldResponseBody `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLayoutFieldAllResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLayoutFieldAllResponse struct{}"
	}

	return strings.Join([]string{"ListLayoutFieldAllResponse", string(data)}, " ")
}
