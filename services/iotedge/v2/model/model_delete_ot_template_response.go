package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteOtTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteOtTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOtTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteOtTemplateResponse", string(data)}, " ")
}
