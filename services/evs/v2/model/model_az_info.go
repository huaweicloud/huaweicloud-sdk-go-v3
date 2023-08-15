package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AzInfo 一个az对象
type AzInfo struct {

	// 可用分区的名字。
	ZoneName string `json:"zoneName"`

	ZoneState *ZoneState `json:"zoneState"`
}

func (o AzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfo struct{}"
	}

	return strings.Join([]string{"AzInfo", string(data)}, " ")
}
