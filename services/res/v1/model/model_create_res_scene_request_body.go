package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateResSceneRequestBody struct {

	// 场景类型： - customize，自定义推荐
	Category string `json:"category" xml:"category"`

	// 数据源id。
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`

	DsConfig *DsConfig `json:"ds_config" xml:"ds_config"`

	// 场景名称，1-64位的字母、数字、下划线、中划线组合。
	SceneName string `json:"scene_name" xml:"scene_name"`

	SpecsConfig *SpecsConfig `json:"specs_config" xml:"specs_config"`

	// 场景类型： - UI，基于用户推荐物品 - UU，基于用户推荐用户 - II，基于物品推荐物品 - IU，基于物品推荐用户
	Type string `json:"type" xml:"type"`

	// 服务类型： - rank，排序服务 - rec，推荐服务
	ServiceType string `json:"service_type" xml:"service_type"`
}

func (o CreateResSceneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResSceneRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResSceneRequestBody", string(data)}, " ")
}
