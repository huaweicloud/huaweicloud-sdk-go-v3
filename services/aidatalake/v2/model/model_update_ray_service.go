package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRayService **参数解释**：更新 RayService 端点配置。支持更新 Serve 应用配置、集群规格、数据预处理配置等，更新时覆盖原有配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type UpdateRayService struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayService。 **默认取值**：不涉及。
	Type *UpdateRayServiceType `json:"type,omitempty"`

	ServeConfig *ServeConfig `json:"serve_config,omitempty"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config,omitempty"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *StoragesConfig `json:"storages_config,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`

	// **参数解释**：模型信息列表（可选）。更新时如果模型没有变化，可以不填此字段；如果需要更新模型，则填写新的模型信息。 **约束限制**：不涉及。
	DataInfos *[]DataInfo `json:"data_infos,omitempty"`
}

func (o UpdateRayService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRayService struct{}"
	}

	return strings.Join([]string{"UpdateRayService", string(data)}, " ")
}

type UpdateRayServiceType struct {
	value string
}

type UpdateRayServiceTypeEnum struct {
	RAY_SERVICE UpdateRayServiceType
}

func GetUpdateRayServiceTypeEnum() UpdateRayServiceTypeEnum {
	return UpdateRayServiceTypeEnum{
		RAY_SERVICE: UpdateRayServiceType{
			value: "RayService",
		},
	}
}

func (c UpdateRayServiceType) Value() string {
	return c.value
}

func (c UpdateRayServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRayServiceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
