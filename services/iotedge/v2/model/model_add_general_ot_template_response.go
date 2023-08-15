package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGeneralOtTemplateResponse Response Object
type AddGeneralOtTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddGeneralOtTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeneralOtTemplateResponse struct{}"
	}

	return strings.Join([]string{"AddGeneralOtTemplateResponse", string(data)}, " ")
}
