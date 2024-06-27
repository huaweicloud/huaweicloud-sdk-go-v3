package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClickHouseDatabaseUserResponse Response Object
type CreateClickHouseDatabaseUserResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClickHouseDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClickHouseDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateClickHouseDatabaseUserResponse", string(data)}, " ")
}
