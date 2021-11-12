package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestartInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}
