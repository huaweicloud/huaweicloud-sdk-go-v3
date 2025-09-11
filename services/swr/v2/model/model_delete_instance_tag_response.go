package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceTagResponse Response Object
type DeleteInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceTagResponse", string(data)}, " ")
}
