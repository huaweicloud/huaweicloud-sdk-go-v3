package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbBorderCross 全域互联带宽跨境属性。
type GcbBorderCross struct {

	// 功能说明：全域互联带宽是否跨境，判断依据：带宽是否涉及从中国大陆到其他国家。 取值范围：True：跨境；False：非跨境
	Bordercross bool `json:"bordercross"`
}

func (o GcbBorderCross) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbBorderCross struct{}"
	}

	return strings.Join([]string{"GcbBorderCross", string(data)}, " ")
}
