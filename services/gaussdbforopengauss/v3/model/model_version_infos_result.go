package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionInfosResult struct {

	// **参数解释**: 引擎版本号。 **取值范围**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: 内核引擎版本号。 **取值范围**: 不涉及
	KernelVersion *string `json:"kernel_version,omitempty"`

	// **参数解释**: 是否为推荐版本。 **取值范围**: - true：推荐版本 - flase: 非推荐版本
	Recommend *bool `json:"recommend,omitempty"`
}

func (o VersionInfosResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionInfosResult struct{}"
	}

	return strings.Join([]string{"VersionInfosResult", string(data)}, " ")
}
