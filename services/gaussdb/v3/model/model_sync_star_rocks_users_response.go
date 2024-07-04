package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncStarRocksUsersResponse Response Object
type SyncStarRocksUsersResponse struct {

	// 开启行列分流是否成功。
	Success        *string `json:"success,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncStarRocksUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncStarRocksUsersResponse struct{}"
	}

	return strings.Join([]string{"SyncStarRocksUsersResponse", string(data)}, " ")
}
