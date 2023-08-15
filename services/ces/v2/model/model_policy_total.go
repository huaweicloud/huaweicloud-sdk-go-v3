package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTotal 告警模板的告警策略总数
type PolicyTotal struct {
}

func (o PolicyTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTotal struct{}"
	}

	return strings.Join([]string{"PolicyTotal", string(data)}, " ")
}
