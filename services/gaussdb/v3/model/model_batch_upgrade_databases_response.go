package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeDatabasesResponse Response Object
type BatchUpgradeDatabasesResponse struct {

	// 接口调用结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpgradeDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabasesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabasesResponse", string(data)}, " ")
}
