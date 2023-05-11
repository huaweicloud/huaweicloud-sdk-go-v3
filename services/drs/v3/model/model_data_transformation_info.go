package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据加工详情
type DataTransformationInfo struct {
	TransformationInfo *TransformationInfo `json:"transformation_info,omitempty"`

	ConfigTransformation *ConfigTransformationVo `json:"config_transformation,omitempty"`

	// 数据加工对象。
	DataTransformationObjectInfos *[]DataTransformationObjectVo `json:"data_transformation_object_infos,omitempty"`
}

func (o DataTransformationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationInfo struct{}"
	}

	return strings.Join([]string{"DataTransformationInfo", string(data)}, " ")
}
