package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateResponse", string(data)}, " ")
}
