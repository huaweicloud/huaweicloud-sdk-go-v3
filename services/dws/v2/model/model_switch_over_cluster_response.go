package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchOverClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchOverClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverClusterResponse struct{}"
	}

	return strings.Join([]string{"SwitchOverClusterResponse", string(data)}, " ")
}
