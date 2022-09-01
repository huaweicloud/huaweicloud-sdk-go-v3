package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisableAttributeResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAttributeResponse struct{}"
	}

	return strings.Join([]string{"DisableAttributeResponse", string(data)}, " ")
}
