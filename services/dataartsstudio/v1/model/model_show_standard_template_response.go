package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardTemplateResponse Response Object
type ShowStandardTemplateResponse struct {
	Data           *ShowStandardTemplateResultData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowStandardTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowStandardTemplateResponse", string(data)}, " ")
}
