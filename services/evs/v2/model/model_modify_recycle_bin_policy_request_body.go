package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRecycleBinPolicyRequestBody 更新回收站策略请求
type ModifyRecycleBinPolicyRequestBody struct {

	// **参数解释** 回收站开关。 **约束限制** 不涉及。 **取值范围** - true：表示打开回收站。 - false：表示关闭回收站。 **默认取值** 不涉及。
	Switch *bool `json:"switch,omitempty"`

	// **参数解释** 回收站门槛时间，云硬盘创建多少天后删除才会放入回收站。 **约束限制** 不涉及。 **取值范围** 1-1000 **默认取值** 7
	ThresholdTime *int32 `json:"threshold_time,omitempty"`

	// **参数解释** 回收站的云硬盘保存期限（天），云硬盘进入回收站后多少天才会被彻底删除。 **约束限制** 不涉及。 **取值范围** 1-365 **默认取值** 7
	KeepTime *int32 `json:"keep_time,omitempty"`
}

func (o ModifyRecycleBinPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRecycleBinPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyRecycleBinPolicyRequestBody", string(data)}, " ")
}
