package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsForMigrateResponse Response Object
type ListAvailableRdsForMigrateResponse struct {

	// 可用后端DN信息。
	DataNodes      *[]AvailableDnInstance `json:"data_nodes,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAvailableRdsForMigrateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsForMigrateResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsForMigrateResponse", string(data)}, " ")
}
