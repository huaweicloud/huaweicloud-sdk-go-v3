package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesNoIndexTablesResponse Response Object
type ListInstancesNoIndexTablesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ListInstancesNoIndexTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesNoIndexTablesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesNoIndexTablesResponse", string(data)}, " ")
}
