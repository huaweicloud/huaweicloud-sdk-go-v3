package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRemediation 补救措施
type AlertRemediation struct {

	// 推荐处理方法
	Recommendation *string `json:"recommendation,omitempty"`

	// 链接，指向该事件的一般修复信息。该URL必须可以从公网访问，不需要提供凭证
	Url *string `json:"url,omitempty"`
}

func (o AlertRemediation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRemediation struct{}"
	}

	return strings.Join([]string{"AlertRemediation", string(data)}, " ")
}
