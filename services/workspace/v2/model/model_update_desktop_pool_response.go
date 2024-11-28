package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopPoolResponse Response Object
type UpdateDesktopPoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesktopPoolResponse", string(data)}, " ")
}
