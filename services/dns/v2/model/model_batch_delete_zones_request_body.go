package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteZonesRequestBody struct {

	// 待删除zone类型，当前仅支持 public 或 private。
	ZoneType string `json:"zone_type"`

	// 待删除Zone ID列表。 最多支持50个。
	ZoneIds []string `json:"zone_ids"`
}

func (o BatchDeleteZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteZonesRequestBody", string(data)}, " ")
}
