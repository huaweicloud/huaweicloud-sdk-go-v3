package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEngine 训练作业的引擎。使用算法管理的算法id或订阅算法subscription_id+item_version_id创建作业时，无需填写。
type JobEngine struct {

	// 训练作业选择的引擎规格ID。engine_id，engine_name+engine_version和image_url方式三选一。
	EngineId *string `json:"engine_id,omitempty"`

	// 训练作业选择的引擎名称。如果已填写engine_id，则此参数无需填写。
	EngineName *string `json:"engine_name,omitempty"`

	// 训练作业选择的引擎版本名称。如果已填写engine_id，则此参数无需填写。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 训练作业选择的自定义镜像地址，地址从swr服务获取。
	ImageUrl *string `json:"image_url,omitempty"`

	// 是否需要安装训练平台指定的 moxing 版本。true为需要。只有填写了engine_name，engine_version，image_url参数时支持该设置。
	InstallSysPackages *bool `json:"install_sys_packages,omitempty"`

	// **参数解释**：SWR企业仓实例ID，使用企业仓镜像时需传入。 **取值范围**：不涉及。
	ImageRepoId *string `json:"image_repo_id,omitempty"`
}

func (o JobEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEngine struct{}"
	}

	return strings.Join([]string{"JobEngine", string(data)}, " ")
}
