package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUrlinReq top_url配置
type TopUrlinReq struct {

	// **参数解释：** 配置开关 **约束限制：** 不涉及 **取值范围：** - true：打开 - false：关闭 **默认取值：** 不涉及
	Enable *bool `json:"enable,omitempty"`
}

func (o TopUrlinReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrlinReq struct{}"
	}

	return strings.Join([]string{"TopUrlinReq", string(data)}, " ")
}
