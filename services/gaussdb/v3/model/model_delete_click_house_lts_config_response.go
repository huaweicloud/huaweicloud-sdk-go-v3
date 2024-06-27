package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseLtsConfigResponse Response Object
type DeleteClickHouseLtsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClickHouseLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseLtsConfigResponse", string(data)}, " ")
}
