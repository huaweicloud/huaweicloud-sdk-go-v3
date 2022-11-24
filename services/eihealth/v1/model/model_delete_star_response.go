package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteStarResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarResponse struct{}"
	}

	return strings.Join([]string{"DeleteStarResponse", string(data)}, " ")
}
