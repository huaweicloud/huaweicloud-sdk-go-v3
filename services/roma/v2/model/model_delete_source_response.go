package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSourceResponse Response Object
type DeleteSourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSourceResponse", string(data)}, " ")
}
