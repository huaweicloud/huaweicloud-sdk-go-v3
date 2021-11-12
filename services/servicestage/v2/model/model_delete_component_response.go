package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentResponse struct{}"
	}

	return strings.Join([]string{"DeleteComponentResponse", string(data)}, " ")
}
