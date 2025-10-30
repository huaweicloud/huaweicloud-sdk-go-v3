package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetInstancesNoIndexTablesResponse Response Object
type GetInstancesNoIndexTablesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GetInstancesNoIndexTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesNoIndexTablesResponse struct{}"
	}

	return strings.Join([]string{"GetInstancesNoIndexTablesResponse", string(data)}, " ")
}
