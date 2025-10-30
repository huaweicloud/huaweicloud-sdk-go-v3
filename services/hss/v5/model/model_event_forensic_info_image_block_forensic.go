package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoImageBlockForensic **参数解释**： 容器镜像阻断取证信息 **取值范围**： 不涉及
type EventForensicInfoImageBlockForensic struct {

	// **参数解释**： 阻断类型 **取值范围**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 阻断原因 **取值范围**： 不涉及
	Msg *string `json:"msg,omitempty"`

	// **参数解释**： 路径 **取值范围**： 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 不涉及
	Image *string `json:"image,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 不涉及
	Version *bool `json:"version,omitempty"`

	// **参数解释**： 阻断结果 **取值范围**： 不涉及
	Result *string `json:"result,omitempty"`

	// **参数解释**： 发生时间 **取值范围**： 不涉及
	Time *string `json:"time,omitempty"`
}

func (o EventForensicInfoImageBlockForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoImageBlockForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoImageBlockForensic", string(data)}, " ")
}
