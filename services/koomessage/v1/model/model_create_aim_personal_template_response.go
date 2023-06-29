package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAimPersonalTemplateResponse Response Object
type CreateAimPersonalTemplateResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *CreateAimPersonalTemplateResponseMode `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o CreateAimPersonalTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimPersonalTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAimPersonalTemplateResponse", string(data)}, " ")
}
