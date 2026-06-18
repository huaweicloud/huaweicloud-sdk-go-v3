package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume volume信息。
type Volume struct {

	// 参数解释： 磁盘大小。单位：GB。 取值范围： 不涉及。
	Size string `json:"size"`

	// 参数解释： 磁盘使用量。单位：GB。 取值范围： 不涉及。
	Used string `json:"used"`

	// 参数解释： 赠送的磁盘大小。单位：GB。 取值范围： 不涉及。
	GiftSize *string `json:"gift_size,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
