package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAllTablesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAllTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllTablesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAllTablesResponse", string(data)}, " ")
}
