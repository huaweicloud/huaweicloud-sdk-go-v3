package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteZonesReq struct {

	// 待删除zone类型，当前仅支持 public 或 private
	ZoneType string `json:"zone_type"`

	// 待删除Zone ID。
	ZoneIds []string `json:"zone_ids"`
}

func (o DeleteZonesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteZonesReq struct{}"
	}

	return strings.Join([]string{"DeleteZonesReq", string(data)}, " ")
}
