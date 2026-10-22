package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HeadGroupSpecV3 **参数解释**：HeadGroup的配置。
type HeadGroupSpecV3 struct {
	Resource *ResourceSpecV3 `json:"resource"`
}

func (o HeadGroupSpecV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeadGroupSpecV3 struct{}"
	}

	return strings.Join([]string{"HeadGroupSpecV3", string(data)}, " ")
}
