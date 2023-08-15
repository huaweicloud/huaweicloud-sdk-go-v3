package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmTemplateResponse Response Object
type UpdateAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateResponse", string(data)}, " ")
}
