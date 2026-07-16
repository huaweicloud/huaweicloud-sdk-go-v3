package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageRegRequest struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。  **默认取值**：X86_64。
	Arch *ImageRegRequestArch `json:"arch,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **约束限制**：不涉及。 **取值范围**：长度限制512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：指定镜像来源，可选项。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE: 用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。  **默认取值**：CUSTOMIZE。
	Origin *ImageRegRequestOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格，默认值CPU、GPU。 枚举值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)。  **约束限制**：不涉及。
	ResourceCategory *[]ImageRegRequestResourceCategory `json:"resource_category,omitempty"`

	// **参数解释**：镜像支持服务类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE: 建议仅在推理部署场景使用。 - TRAIN: 建议仅在训练任务场景使用。 - DEV: 建议仅在开发调测场景使用。 - UNKNOWN: 未明确设置的镜像支持的服务类型。  **默认取值**：UNKNOWN。
	ServiceType *ImageRegRequestServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像支持的服务，默认值NOTEBOOK、SSH。枚举值如下: - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。  **约束限制**：不涉及。
	Services *[]ImageRegRequestServices `json:"services,omitempty"`

	// **参数解释**：SWR镜像地址。 **约束限制**：不涉及。 **取值范围**：长度最长为2048个字符，最短为16个字符，地址格式为：[仓库地址[:端口]]/[命名空间]/[镜像名称]:[标签]。 **默认取值**：不涉及。
	SwrPath string `json:"swr_path"`

	// **参数解释**：镜像可见度。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC: 所有用户可以根据image_id来进行只读使用。  **默认取值**：PRIVATE。
	Visibility *ImageRegRequestVisibility `json:"visibility,omitempty"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：0或32位仅包含字符0-9或小写字母a-z的字符串。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： -ASCEND_SNT9：昇腾910芯片。 -ASCEND_SNT9B：昇腾910B芯片。 -ASCEND_SNT3：昇腾310芯片。  **默认取值**：不涉及。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：该镜像所属镜像组对应的标签。 **约束限制**：最大支持20个标签。 **取值范围**：key值最大支持长度128，value值最大支持255。 **默认取值**：不涉及。
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释**：企业版SWR仓库ID。 **参数约束**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`

	// **参数解释**：镜像指导。 **参数约束**：不涉及。 **取值范围**：字符串长度限制为3000个字符。 **默认取值**：不涉及。
	ReadMe *string `json:"read_me,omitempty"`
}

func (o ImageRegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageRegRequest struct{}"
	}

	return strings.Join([]string{"ImageRegRequest", string(data)}, " ")
}

type ImageRegRequestArch struct {
	value string
}

type ImageRegRequestArchEnum struct {
	AARCH64 ImageRegRequestArch
	X86_64  ImageRegRequestArch
}

func GetImageRegRequestArchEnum() ImageRegRequestArchEnum {
	return ImageRegRequestArchEnum{
		AARCH64: ImageRegRequestArch{
			value: "AARCH64",
		},
		X86_64: ImageRegRequestArch{
			value: "X86_64",
		},
	}
}

func (c ImageRegRequestArch) Value() string {
	return c.value
}

func (c ImageRegRequestArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestArch) UnmarshalJSON(b []byte) error {
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

type ImageRegRequestOrigin struct {
	value string
}

type ImageRegRequestOriginEnum struct {
	CUSTOMIZE  ImageRegRequestOrigin
	IMAGE_SAVE ImageRegRequestOrigin
}

func GetImageRegRequestOriginEnum() ImageRegRequestOriginEnum {
	return ImageRegRequestOriginEnum{
		CUSTOMIZE: ImageRegRequestOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: ImageRegRequestOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c ImageRegRequestOrigin) Value() string {
	return c.value
}

func (c ImageRegRequestOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestOrigin) UnmarshalJSON(b []byte) error {
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

type ImageRegRequestResourceCategory struct {
	value string
}

type ImageRegRequestResourceCategoryEnum struct {
	ASCEND ImageRegRequestResourceCategory
	CPU    ImageRegRequestResourceCategory
	GPU    ImageRegRequestResourceCategory
}

func GetImageRegRequestResourceCategoryEnum() ImageRegRequestResourceCategoryEnum {
	return ImageRegRequestResourceCategoryEnum{
		ASCEND: ImageRegRequestResourceCategory{
			value: "ASCEND",
		},
		CPU: ImageRegRequestResourceCategory{
			value: "CPU",
		},
		GPU: ImageRegRequestResourceCategory{
			value: "GPU",
		},
	}
}

func (c ImageRegRequestResourceCategory) Value() string {
	return c.value
}

func (c ImageRegRequestResourceCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestResourceCategory) UnmarshalJSON(b []byte) error {
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

type ImageRegRequestServiceType struct {
	value string
}

type ImageRegRequestServiceTypeEnum struct {
	COMMON    ImageRegRequestServiceType
	DEV       ImageRegRequestServiceType
	INFERENCE ImageRegRequestServiceType
	TRAIN     ImageRegRequestServiceType
	UNKNOWN   ImageRegRequestServiceType
}

func GetImageRegRequestServiceTypeEnum() ImageRegRequestServiceTypeEnum {
	return ImageRegRequestServiceTypeEnum{
		COMMON: ImageRegRequestServiceType{
			value: "COMMON",
		},
		DEV: ImageRegRequestServiceType{
			value: "DEV",
		},
		INFERENCE: ImageRegRequestServiceType{
			value: "INFERENCE",
		},
		TRAIN: ImageRegRequestServiceType{
			value: "TRAIN",
		},
		UNKNOWN: ImageRegRequestServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c ImageRegRequestServiceType) Value() string {
	return c.value
}

func (c ImageRegRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestServiceType) UnmarshalJSON(b []byte) error {
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

type ImageRegRequestServices struct {
	value string
}

type ImageRegRequestServicesEnum struct {
	AI_FLOW      ImageRegRequestServices
	MA_STUDIO    ImageRegRequestServices
	NOTEBOOK     ImageRegRequestServices
	SSH          ImageRegRequestServices
	TENSOR_BOARD ImageRegRequestServices
	WEB_IDE      ImageRegRequestServices
}

func GetImageRegRequestServicesEnum() ImageRegRequestServicesEnum {
	return ImageRegRequestServicesEnum{
		AI_FLOW: ImageRegRequestServices{
			value: "AI_FLOW",
		},
		MA_STUDIO: ImageRegRequestServices{
			value: "MA_STUDIO",
		},
		NOTEBOOK: ImageRegRequestServices{
			value: "NOTEBOOK",
		},
		SSH: ImageRegRequestServices{
			value: "SSH",
		},
		TENSOR_BOARD: ImageRegRequestServices{
			value: "TENSOR_BOARD",
		},
		WEB_IDE: ImageRegRequestServices{
			value: "WEB_IDE",
		},
	}
}

func (c ImageRegRequestServices) Value() string {
	return c.value
}

func (c ImageRegRequestServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestServices) UnmarshalJSON(b []byte) error {
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

type ImageRegRequestVisibility struct {
	value string
}

type ImageRegRequestVisibilityEnum struct {
	HIDDEN  ImageRegRequestVisibility
	PRIVATE ImageRegRequestVisibility
	PUBLIC  ImageRegRequestVisibility
}

func GetImageRegRequestVisibilityEnum() ImageRegRequestVisibilityEnum {
	return ImageRegRequestVisibilityEnum{
		HIDDEN: ImageRegRequestVisibility{
			value: "HIDDEN",
		},
		PRIVATE: ImageRegRequestVisibility{
			value: "PRIVATE",
		},
		PUBLIC: ImageRegRequestVisibility{
			value: "PUBLIC",
		},
	}
}

func (c ImageRegRequestVisibility) Value() string {
	return c.value
}

func (c ImageRegRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRegRequestVisibility) UnmarshalJSON(b []byte) error {
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
