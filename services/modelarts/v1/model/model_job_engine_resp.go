package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEngineResp 训练作业的引擎。使用算法管理的算法id或订阅算法subscription_id+item_version_id创建作业时，无需填写。
type JobEngineResp struct {

	// **参数解释**：训练作业选择的引擎规格ID。 **取值范围**：不涉及。
	EngineId *string `json:"engine_id,omitempty"`

	// **参数解释**：训练作业选择的引擎规格名称。 **取值范围**：不涉及。
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**：训练作业选择的引擎规格版本。 **取值范围**：不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**：训练作业选择的自定义镜像地址，地址从swr服务获取。 **取值范围**：不涉及。
	ImageUrl *string `json:"image_url,omitempty"`

	// **参数解释**：是否需要安装训练平台指定的 moxing 版本。 **取值范围**： - true：需要 - false：不需要
	InstallSysPackages *bool `json:"install_sys_packages,omitempty"`

	// **参数解释**：SWR企业仓实例ID，使用企业仓镜像时需传入。 **取值范围**：不涉及。
	ImageRepoId *string `json:"image_repo_id,omitempty"`
}

func (o JobEngineResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEngineResp struct{}"
	}

	return strings.Join([]string{"JobEngineResp", string(data)}, " ")
}
