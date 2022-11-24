package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNodeEncryptdatasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNodeEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"CreateNodeEncryptdatasResponse", string(data)}, " ")
}
