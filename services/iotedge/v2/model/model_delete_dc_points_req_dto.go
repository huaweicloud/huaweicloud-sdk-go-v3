package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointsReqDto 批量删除点位表配置请求结构体
type DeleteDcPointsReqDto struct {
	Points []string `json:"points"`
}

func (o DeleteDcPointsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointsReqDto struct{}"
	}

	return strings.Join([]string{"DeleteDcPointsReqDto", string(data)}, " ")
}
