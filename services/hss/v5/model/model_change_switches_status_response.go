package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSwitchesStatusResponse Response Object
type ChangeSwitchesStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeSwitchesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSwitchesStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeSwitchesStatusResponse", string(data)}, " ")
}
