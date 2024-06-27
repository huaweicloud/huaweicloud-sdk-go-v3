package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestcaseByPageResponse Response Object
type ShowTestcaseByPageResponse struct {
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoTestCase `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTestcaseByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestcaseByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowTestcaseByPageResponse", string(data)}, " ")
}
