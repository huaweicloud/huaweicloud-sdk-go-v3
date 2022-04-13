package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisableKeyRotationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableKeyRotationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyRotationResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyRotationResponse", string(data)}, " ")
}
