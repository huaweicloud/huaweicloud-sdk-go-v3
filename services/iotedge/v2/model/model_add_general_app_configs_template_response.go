package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGeneralAppConfigsTemplateResponse Response Object
type AddGeneralAppConfigsTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddGeneralAppConfigsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeneralAppConfigsTemplateResponse struct{}"
	}

	return strings.Join([]string{"AddGeneralAppConfigsTemplateResponse", string(data)}, " ")
}
