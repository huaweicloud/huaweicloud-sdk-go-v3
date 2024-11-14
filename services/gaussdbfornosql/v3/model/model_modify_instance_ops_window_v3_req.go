package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyInstanceOpsWindowV3Req struct {

	// 参数解释： 开始时间，格式必须为HH:MM且有效，时间为UTC时间，只能为整点。结束时间默认与开始时间相隔四个小时。 约束限制： 不涉及。
	StartTime string `json:"start_time"`
}

func (o ModifyInstanceOpsWindowV3Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceOpsWindowV3Req struct{}"
	}

	return strings.Join([]string{"ModifyInstanceOpsWindowV3Req", string(data)}, " ")
}
