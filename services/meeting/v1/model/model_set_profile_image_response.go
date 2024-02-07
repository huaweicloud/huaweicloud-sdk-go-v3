package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProfileImageResponse Response Object
type SetProfileImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetProfileImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProfileImageResponse struct{}"
	}

	return strings.Join([]string{"SetProfileImageResponse", string(data)}, " ")
}
