package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserSettingResponse Response Object
type ShowUserSettingResponse struct {
	Operation *Operation `json:"operation,omitempty"`

	Settings       *UserSettingDto `json:"settings,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowUserSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowUserSettingResponse", string(data)}, " ")
}
