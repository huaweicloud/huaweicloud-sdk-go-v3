package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopResponse Response Object
type DeleteDesktopResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesktopResponse", string(data)}, " ")
}
