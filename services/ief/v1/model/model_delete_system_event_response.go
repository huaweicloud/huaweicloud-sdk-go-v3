package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSystemEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSystemEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSystemEventResponse struct{}"
	}

	return strings.Join([]string{"DeleteSystemEventResponse", string(data)}, " ")
}
