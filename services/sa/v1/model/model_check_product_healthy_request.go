package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckProductHealthyRequest struct {
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *ProductInfo `json:"body,omitempty" xml:"body"`
}

func (o CheckProductHealthyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProductHealthyRequest struct{}"
	}

	return strings.Join([]string{"CheckProductHealthyRequest", string(data)}, " ")
}
