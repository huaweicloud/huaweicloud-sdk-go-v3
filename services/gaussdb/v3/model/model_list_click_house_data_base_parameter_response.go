package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseDataBaseParameterResponse Response Object
type ListClickHouseDataBaseParameterResponse struct {

	// 库参数信息。
	DbParameters   *[]ChDatabaseParameterInfo `json:"db_parameters,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListClickHouseDataBaseParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseDataBaseParameterResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseDataBaseParameterResponse", string(data)}, " ")
}
