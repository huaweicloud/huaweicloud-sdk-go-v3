package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProfileResponse Response Object
type DeleteProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProfileResponse struct{}"
	}

	return strings.Join([]string{"DeleteProfileResponse", string(data)}, " ")
}
