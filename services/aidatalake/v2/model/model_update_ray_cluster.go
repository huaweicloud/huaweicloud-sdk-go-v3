package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRayCluster **参数解释**：Ray集群配置。
type UpdateRayCluster struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayCluster。 **默认取值**：不涉及。
	Type *UpdateRayClusterType `json:"type,omitempty"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config,omitempty"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *StoragesConfig `json:"storages_config,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`
}

func (o UpdateRayCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRayCluster struct{}"
	}

	return strings.Join([]string{"UpdateRayCluster", string(data)}, " ")
}

type UpdateRayClusterType struct {
	value string
}

type UpdateRayClusterTypeEnum struct {
	RAY_CLUSTER UpdateRayClusterType
}

func GetUpdateRayClusterTypeEnum() UpdateRayClusterTypeEnum {
	return UpdateRayClusterTypeEnum{
		RAY_CLUSTER: UpdateRayClusterType{
			value: "RayCluster",
		},
	}
}

func (c UpdateRayClusterType) Value() string {
	return c.value
}

func (c UpdateRayClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRayClusterType) UnmarshalJSON(b []byte) error {
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
