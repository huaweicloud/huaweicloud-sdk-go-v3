package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbHotfixInfoResult struct {

	// **参数解释**： 热补丁版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 热补丁的创建时间，UNIX时间戳格式，单位是毫秒。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 可升级该补丁的实例部署形态信息。
	DeployModes *[]HotfixDeployMode `json:"deploy_modes,omitempty"`
}

func (o DbHotfixInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbHotfixInfoResult struct{}"
	}

	return strings.Join([]string{"DbHotfixInfoResult", string(data)}, " ")
}
