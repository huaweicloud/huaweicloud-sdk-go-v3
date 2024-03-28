package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbShowLocalArea struct {

	// 功能说明：本端接入点的中英文名。通过HEADER里面的x-language控制，默认英文，zh-cn返回中文。
	LocalArea *string `json:"local_area,omitempty"`
}

func (o GcbShowLocalArea) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbShowLocalArea struct{}"
	}

	return strings.Join([]string{"GcbShowLocalArea", string(data)}, " ")
}
