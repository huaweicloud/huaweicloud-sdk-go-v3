package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolResponse Response Object
type DeleteDesktopPoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolResponse", string(data)}, " ")
}
