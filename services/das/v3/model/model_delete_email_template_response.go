package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEmailTemplateResponse Response Object
type DeleteEmailTemplateResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteEmailTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEmailTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteEmailTemplateResponse", string(data)}, " ")
}
