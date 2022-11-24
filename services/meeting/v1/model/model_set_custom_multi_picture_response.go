package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetCustomMultiPictureResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetCustomMultiPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCustomMultiPictureResponse struct{}"
	}

	return strings.Join([]string{"SetCustomMultiPictureResponse", string(data)}, " ")
}
