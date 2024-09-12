package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointsFailedDetail 批量删除点位表配置响应结构体
type DeleteDcPointsFailedDetail struct {

	// 点位删除失败错误码
	ErrorCode string `json:"error_code"`

	// 点位删除失败错误详情
	ErrorMsg string `json:"error_msg"`

	// 点位id
	PointId string `json:"point_id"`
}

func (o DeleteDcPointsFailedDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointsFailedDetail struct{}"
	}

	return strings.Join([]string{"DeleteDcPointsFailedDetail", string(data)}, " ")
}
