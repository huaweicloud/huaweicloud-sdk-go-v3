package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResSceneRequestBody This is a auto create Body Object
type UpdateResSceneRequestBody struct {

	// 场景类型： - customize，自定义推荐 - popularity，热门推荐 - relation，关联推荐 - personalization，猜你喜欢
	Category string `json:"category"`

	// 数据源id，字母、数字、下划线、减号组合32位。
	DatasourceId string `json:"datasource_id"`

	DsConfig *DsConfig `json:"ds_config"`

	// 场景名称，1-64位的字母、数字、下划线、中划线组合。
	SceneName string `json:"scene_name"`

	SpecsConfig *SpecsConfig `json:"specs_config"`

	// 场景类型： - UI，基于用户推荐物品 - UU，基于用户推荐用户 - II，基于物品推荐物品 - IU，基于物品推荐用户
	Type string `json:"type"`

	// 服务类型： - rank，排序服务 - rec，推荐服务
	ServiceType string `json:"service_type"`
}

func (o UpdateResSceneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResSceneRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResSceneRequestBody", string(data)}, " ")
}
