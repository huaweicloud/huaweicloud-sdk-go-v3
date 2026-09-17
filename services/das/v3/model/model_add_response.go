package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddResponse Response Object
type AddResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResponse struct{}"
	}

	return strings.Join([]string{"AddResponse", string(data)}, " ")
}
