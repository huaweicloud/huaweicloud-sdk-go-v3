package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCustomMultiPictureResponse Response Object
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
