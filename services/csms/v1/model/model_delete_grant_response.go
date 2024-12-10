package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGrantResponse Response Object
type DeleteGrantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGrantResponse struct{}"
	}

	return strings.Join([]string{"DeleteGrantResponse", string(data)}, " ")
}
