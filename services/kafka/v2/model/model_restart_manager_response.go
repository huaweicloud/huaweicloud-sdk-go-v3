package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestartManagerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartManagerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartManagerResponse struct{}"
	}

	return strings.Join([]string{"RestartManagerResponse", string(data)}, " ")
}
