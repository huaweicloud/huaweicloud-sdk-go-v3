package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRayCluster **参数解释**：Ray集群配置。
type CreateRayCluster struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayCluster。 **默认取值**：不涉及。
	Type *CreateRayClusterType `json:"type,omitempty"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *CreateStoragesConfig `json:"storages_config,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image"`

	HistoryServerConfig *CreateHistoryServerConfig `json:"history_server_config,omitempty"`
}

func (o CreateRayCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRayCluster struct{}"
	}

	return strings.Join([]string{"CreateRayCluster", string(data)}, " ")
}

type CreateRayClusterType struct {
	value string
}

type CreateRayClusterTypeEnum struct {
	RAY_CLUSTER CreateRayClusterType
}

func GetCreateRayClusterTypeEnum() CreateRayClusterTypeEnum {
	return CreateRayClusterTypeEnum{
		RAY_CLUSTER: CreateRayClusterType{
			value: "RayCluster",
		},
	}
}

func (c CreateRayClusterType) Value() string {
	return c.value
}

func (c CreateRayClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRayClusterType) UnmarshalJSON(b []byte) error {
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
