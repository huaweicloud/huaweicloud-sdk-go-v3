package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotebookUpdateRequest struct {

	// **参数解释**：支持更新实例描述信息。 **约束限制**：不涉及。 **取值范围**：长度限制为512字符，且不能包含字符&<>\"'/。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：仅在本地IDE（如PyCharm、VS Code）或SSH客户端接入Notebook。 **约束限制**：仅在本地IDE（如PyCharm、VS Code）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
	Endpoints *[]EndpointsReq `json:"endpoints,omitempty"`

	// **参数解释**：支持变更实例的机器规格。支持变更的规格可以通过本章节的[查询支持可切换规格列表](ShowSwitchableFlavors.xml)的API获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	CustomSpec *NotebookCustomSpec `json:"custom_spec,omitempty"`

	// **参数解释**：支持更新镜像ID，镜像ID参考[查询支持的镜像列表](ListImage.xml)获取。 **约束限制**：不涉及。 **取值范围**：调用[查询支持的镜像列表](ListImage.xml)接口获取的合法镜像ID列表。 **默认取值**：不涉及。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：支持更新实例名称。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线和下划线，名称可重复。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：EVS实例支持动态扩充的容量，单位GB。只允许扩容，不允许缩容。 **约束限制**：不涉及。 **取值范围**：最大允许扩容至4096。 **默认取值**：不涉及。
	StorageNewSize *int32 `json:"storage_new_size,omitempty"`

	Hooks *CustomHooks `json:"hooks,omitempty"`

	Affinity *AffinityType `json:"affinity,omitempty"`

	// **参数解释**：DEW存储的用户AKSK凭据名称。 **参数约束**：当category为OBS时必填，仅支持大小写字母、数字、中划线、下划线，长度 1-64 字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DewSecretName *string `json:"dew_secret_name,omitempty"`

	// **参数解释**：扩展存储信息。 **约束限制**：不涉及。
	DataVolumes *[]VolumeMountRequest `json:"data_volumes,omitempty"`
}

func (o NotebookUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookUpdateRequest struct{}"
	}

	return strings.Join([]string{"NotebookUpdateRequest", string(data)}, " ")
}
