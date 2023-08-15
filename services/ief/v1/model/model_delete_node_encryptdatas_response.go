package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeEncryptdatasResponse Response Object
type DeleteNodeEncryptdatasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNodeEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodeEncryptdatasResponse", string(data)}, " ")
}
