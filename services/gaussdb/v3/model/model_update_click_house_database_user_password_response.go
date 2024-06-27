package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseDatabaseUserPasswordResponse Response Object
type UpdateClickHouseDatabaseUserPasswordResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClickHouseDatabaseUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseDatabaseUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseDatabaseUserPasswordResponse", string(data)}, " ")
}
