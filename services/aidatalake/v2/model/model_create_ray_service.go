package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRayService **参数解释**：RayService（推理服务）配置。在 RayCluster 能力基础上，新增模型信息（data_infos，仅需 OBS 输入路径）与 Serve 应用配置（serve_config）。 **约束限制**：serve_config、cluster_config、storages_config、image 为必填字段。 **取值范围**：不涉及。 **默认取值**：不涉及。
type CreateRayService struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayService。 **默认取值**：不涉及。
	Type *CreateRayServiceType `json:"type,omitempty"`

	ServeConfig *ServeConfig `json:"serve_config"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *CreateStoragesConfig `json:"storages_config"`

	Image *V2ImageBriefInfoAdmin `json:"image"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`

	// **参数解释**：模型信息列表（可选）。模型可以打包在镜像中（通过 working_dir 配置项指定），也可以通过此字段上传模型地址。创建时如果模型已预置在镜像中，可省略此字段；更新时如果模型没有变化，也可以不填。 **约束限制**：不涉及。
	DataInfos *[]DataInfo `json:"data_infos,omitempty"`
}

func (o CreateRayService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRayService struct{}"
	}

	return strings.Join([]string{"CreateRayService", string(data)}, " ")
}

type CreateRayServiceType struct {
	value string
}

type CreateRayServiceTypeEnum struct {
	RAY_SERVICE CreateRayServiceType
}

func GetCreateRayServiceTypeEnum() CreateRayServiceTypeEnum {
	return CreateRayServiceTypeEnum{
		RAY_SERVICE: CreateRayServiceType{
			value: "RayService",
		},
	}
}

func (c CreateRayServiceType) Value() string {
	return c.value
}

func (c CreateRayServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRayServiceType) UnmarshalJSON(b []byte) error {
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
