package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDbParametersResponse Response Object
type ListStarRocksDbParametersResponse struct {

	// 库参数信息。
	DbParameters   *[]DbParameter `json:"db_parameters,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStarRocksDbParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDbParametersResponse struct{}"
	}

	return strings.Join([]string{"ListStarRocksDbParametersResponse", string(data)}, " ")
}
