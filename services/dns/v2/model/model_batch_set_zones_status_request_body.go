package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetZonesStatusRequestBody struct {

	// 待设置域名状态，支持DISABLE或ENABLE。
	Status string `json:"status"`

	// 待设置域名ID列表。 最多支持50个。
	ZoneIds []string `json:"zone_ids"`
}

func (o BatchSetZonesStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetZonesStatusRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetZonesStatusRequestBody", string(data)}, " ")
}
