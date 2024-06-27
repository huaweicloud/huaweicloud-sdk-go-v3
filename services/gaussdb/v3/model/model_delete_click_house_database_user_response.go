package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDatabaseUserResponse Response Object
type DeleteClickHouseDatabaseUserResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClickHouseDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDatabaseUserResponse", string(data)}, " ")
}
