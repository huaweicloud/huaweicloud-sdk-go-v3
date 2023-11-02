package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDbPositionReq 采集数据库位点信息
type QueryDbPositionReq struct {

	// 重置位点时间, 使用UTC时间 示例：2023-09-19 15:00:00，UTC时间是2023-09-19T07:00:00Z
	ResetPositionTime string `json:"reset_position_time"`
}

func (o QueryDbPositionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDbPositionReq struct{}"
	}

	return strings.Join([]string{"QueryDbPositionReq", string(data)}, " ")
}
