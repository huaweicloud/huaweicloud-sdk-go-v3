package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAttributeResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttributeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAttributeResponse", string(data)}, " ")
}
