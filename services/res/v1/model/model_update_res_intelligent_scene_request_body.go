package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateResIntelligentSceneRequestBody struct {

	// 场景类型： - customize，自定义场景 - intelligent，智能场景
	Category string `json:"category" xml:"category"`

	// 数据源id，字母、数字、下划线、减号组合32位。
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`

	SpecsConfig *SpecsConfig `json:"specs_config" xml:"specs_config"`

	// 调度信息。
	Schedule *string `json:"schedule,omitempty" xml:"schedule"`

	JobConfigs *JobConfig `json:"job_configs" xml:"job_configs"`

	// 场景名称，1-64位字母、数字、下划线、中划线组合。
	SceneName string `json:"scene_name" xml:"scene_name"`
}

func (o UpdateResIntelligentSceneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResIntelligentSceneRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResIntelligentSceneRequestBody", string(data)}, " ")
}
