package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseSchemaResponse Response Object
type DeleteDatabaseSchemaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseSchemaResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseSchemaResponse", string(data)}, " ")
}
