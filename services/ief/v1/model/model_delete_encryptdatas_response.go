package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEncryptdatasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"DeleteEncryptdatasResponse", string(data)}, " ")
}
