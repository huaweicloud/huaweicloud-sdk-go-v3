package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEmailTemplateResponse Response Object
type UpdateEmailTemplateResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateEmailTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEmailTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateEmailTemplateResponse", string(data)}, " ")
}
