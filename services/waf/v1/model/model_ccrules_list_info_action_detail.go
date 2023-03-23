package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 阻断页面信息。当防护动作（category）选择阻断（block）或者动态阻断（dynamic_block）时，需要设置返回的阻断页面。  - 如果需要返回的阻断页面为系统默认的阻断页面，不需要传该参数。  - 如果用户想防护自定义的阻断页面，可以通过此参数设置
type CcrulesListInfoActionDetail struct {
	Response *CcrulesListInfoActionDetailResponse `json:"response,omitempty"`
}

func (o CcrulesListInfoActionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcrulesListInfoActionDetail struct{}"
	}

	return strings.Join([]string{"CcrulesListInfoActionDetail", string(data)}, " ")
}
