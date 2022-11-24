package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEventSchemaVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEventSchemaVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSchemaVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventSchemaVersionResponse", string(data)}, " ")
}
