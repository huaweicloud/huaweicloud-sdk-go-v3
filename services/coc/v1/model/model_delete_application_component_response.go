package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationComponentResponse Response Object
type DeleteApplicationComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationComponentResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationComponentResponse", string(data)}, " ")
}
