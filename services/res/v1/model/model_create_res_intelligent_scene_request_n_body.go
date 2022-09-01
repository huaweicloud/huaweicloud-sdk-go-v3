package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateResIntelligentSceneRequestNBody struct {

	// 场景类型： - popularity，热门推荐 - relation，关联推荐 - personalization，猜你喜欢
	Category string `json:"category" xml:"category"`

	// 数据源id。
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`

	// 场景名称:字母、数字、下划线、中划线组合。
	SceneName string `json:"scene_name" xml:"scene_name"`

	SpecsConfig *SpecsConfig `json:"specs_config" xml:"specs_config"`

	// 调度信息。
	Schedule *string `json:"schedule,omitempty" xml:"schedule"`

	JobConfigs *JobConfig `json:"job_configs" xml:"job_configs"`
}

func (o CreateResIntelligentSceneRequestNBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResIntelligentSceneRequestNBody struct{}"
	}

	return strings.Join([]string{"CreateResIntelligentSceneRequestNBody", string(data)}, " ")
}
