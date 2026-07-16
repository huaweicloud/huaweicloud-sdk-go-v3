package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotebookCreateRequest struct {

	// **参数解释**：实例描述信息。 **约束限制**：不涉及。 **取值范围**：长度限制为512字符，且不能包含字符&<>\"'/。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：仅在本地IDE（如PyCharm、VS Code）或SSH客户端接入Notebook。 **约束限制**：仅在本地IDE（如PyCharm、VS Code）或SSH客户端，通过SSH远程接入Notebook实例时需要的相关配置。
	Endpoints *[]EndpointsReq `json:"endpoints,omitempty"`

	// **参数解释**：实例类别。 **约束限制**：不涉及。 **取值范围**： - DEFAULT：CodeLab免费规格实例，每个用户最多只能创建一个。 - NOTEBOOK：计费规格实例。  **默认取值**：NOTEBOOK。
	Feature *NotebookCreateRequestFeature `json:"feature,omitempty"`

	// **参数解释**：实例的机器规格。如下规格仅供参考，实际支持的规格以具体区域为准。 - modelarts.vm.cpu.2u：Intel CPU通用规格，用于快速数据探索和实验。 - modelarts.vm.cpu.8u：Intel CPU算力增强型，适用于密集计算场景下运算。  **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	CustomSpec *NotebookCustomSpec `json:"custom_spec,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。镜像的ID可通过调用[[查询支持的镜像列表](https://support.huaweicloud.com/api-modelarts/ListImage.html)](tag:hc)[[查询支持的镜像列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListImage.html)](tag:hk)接口获取。 **约束限制**：不涉及。 **取值范围**：调用[[查询支持的镜像列表](https://support.huaweicloud.com/api-modelarts/ListImage.html)](tag:hc)[[查询支持的镜像列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListImage.html)](tag:hk)接口获取的合法镜像ID列表。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：实例名称。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线和下划线，名称可重复。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：专属资源池ID，若需要指定专属资源池创建实例时必填。专属资源池ID可通过[[查询资源池列表](https://support.huaweicloud.com/api-modelarts/ListPools.html)](tag:hc)[[查询资源池列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListPools.html)](tag:hk)接口获取。 **约束限制**：不涉及。 **取值范围**：调用[[查询资源池列表](https://support.huaweicloud.com/api-modelarts/ListPools.html)](tag:hc)[[查询资源池列表](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListPools.html)](tag:hk)接口获取的合法资源池ID列表。 **默认取值**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	Volume *VolumeMountRequest `json:"volume"`

	// **参数解释**：工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：0或32位仅包含字符0-9或小写字母a-z的字符串。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Hooks *CustomHooks `json:"hooks,omitempty"`

	Lease *LeaseReq `json:"lease,omitempty"`

	Affinity *AffinityType `json:"affinity,omitempty"`

	RunUser *RunUserRequest `json:"run_user,omitempty"`

	// **参数解释**：实例存储配置。 **约束限制**：不涉及。
	DataVolumes *[]VolumeMountRequest `json:"data_volumes,omitempty"`

	UserVpc *UserVpcRequest `json:"user_vpc,omitempty"`

	// **参数解释**：定时停止，以当前时刻为起点，运行的时长（到期后自动停止）。单位：毫秒。 **约束限制**：不涉及。 **取值范围**：3600000-259200000。 **默认取值**：3600000。
	Duration *int32 `json:"duration,omitempty"`
}

func (o NotebookCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookCreateRequest struct{}"
	}

	return strings.Join([]string{"NotebookCreateRequest", string(data)}, " ")
}

type NotebookCreateRequestFeature struct {
	value string
}

type NotebookCreateRequestFeatureEnum struct {
	DEFAULT  NotebookCreateRequestFeature
	NOTEBOOK NotebookCreateRequestFeature
}

func GetNotebookCreateRequestFeatureEnum() NotebookCreateRequestFeatureEnum {
	return NotebookCreateRequestFeatureEnum{
		DEFAULT: NotebookCreateRequestFeature{
			value: "DEFAULT",
		},
		NOTEBOOK: NotebookCreateRequestFeature{
			value: "NOTEBOOK",
		},
	}
}

func (c NotebookCreateRequestFeature) Value() string {
	return c.value
}

func (c NotebookCreateRequestFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookCreateRequestFeature) UnmarshalJSON(b []byte) error {
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
