package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Remediation struct {

	// 处理建议，最长512个字符。
	Recommendation string `json:"recommendation" xml:"recommendation"`

	// 链接，指向该事件的一般修复信息。该URL必须可以从公网访问，不需要提供凭证。
	Url *string `json:"url,omitempty" xml:"url"`
}

func (o Remediation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Remediation struct{}"
	}

	return strings.Join([]string{"Remediation", string(data)}, " ")
}
