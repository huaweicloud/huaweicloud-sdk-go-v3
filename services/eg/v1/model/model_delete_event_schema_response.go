package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventSchemaResponse Response Object
type DeleteEventSchemaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEventSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSchemaResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventSchemaResponse", string(data)}, " ")
}
