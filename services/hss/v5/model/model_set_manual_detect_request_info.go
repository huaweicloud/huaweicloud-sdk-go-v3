package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetManualDetectRequestInfo struct {

	// **参数解释**： 手动检测的功能 **约束限制**： 不涉及 **取值范围**： - pwd：弱口令检测  **默认取值**： 不涉及
	Type string `json:"type"`
}

func (o SetManualDetectRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetManualDetectRequestInfo struct{}"
	}

	return strings.Join([]string{"SetManualDetectRequestInfo", string(data)}, " ")
}
