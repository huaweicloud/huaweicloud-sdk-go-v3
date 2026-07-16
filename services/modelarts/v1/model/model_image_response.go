package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageResponse 实例镜像信息。
type ImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *ImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息，长度限制512个字符。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。元素为枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]ImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE: 用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *ImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。元素为枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]ImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE: 建议仅在推理部署场景使用。 - TRAIN: 建议仅在训练任务场景使用。 - DEV: 建议仅在开发调测场景使用。 - UNKNOWN: 未明确设置的镜像支持的服务类型。
	ServiceType *ImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *ImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *ImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC: 所有用户可以根据image_id来进行只读使用。
	Visibility *ImageResponseVisibility `json:"visibility,omitempty"`

	// **参数解释**：工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：镜像的资源类型。 **取值范围**：枚举类型，取值如下： -ASCEND_SNT9：昇腾910芯片。 -ASCEND_SNT9B：昇腾910B芯片。 -ASCEND_SNT3：昇腾310芯片。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：SWR企业仓库ID。未使用SWR企业仓时该字段为null。 **约束限制**：不涉及。 **取值范围**：128位UUID。 **默认取值**：null。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`

	// **参数解释**：镜像展示名称，仅预置镜像具备该字段。
	ShowName *string `json:"show_name,omitempty"`

	// **参数解释**：镜像展示版本号，仅预置镜像具备该字段。
	ShowTag *string `json:"show_tag,omitempty"`

	// **参数解释**：镜像标签。
	Tags *[]TmsTagResponse `json:"tags,omitempty"`
}

func (o ImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageResponse struct{}"
	}

	return strings.Join([]string{"ImageResponse", string(data)}, " ")
}

type ImageResponseArch struct {
	value string
}

type ImageResponseArchEnum struct {
	AARCH64 ImageResponseArch
	X86_64  ImageResponseArch
}

func GetImageResponseArchEnum() ImageResponseArchEnum {
	return ImageResponseArchEnum{
		AARCH64: ImageResponseArch{
			value: "AARCH64",
		},
		X86_64: ImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c ImageResponseArch) Value() string {
	return c.value
}

func (c ImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseArch) UnmarshalJSON(b []byte) error {
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

type ImageResponseDevServices struct {
	value string
}

type ImageResponseDevServicesEnum struct {
	NOTEBOOK ImageResponseDevServices
	SSH      ImageResponseDevServices
}

func GetImageResponseDevServicesEnum() ImageResponseDevServicesEnum {
	return ImageResponseDevServicesEnum{
		NOTEBOOK: ImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: ImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c ImageResponseDevServices) Value() string {
	return c.value
}

func (c ImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type ImageResponseOrigin struct {
	value string
}

type ImageResponseOriginEnum struct {
	CUSTOMIZE  ImageResponseOrigin
	IMAGE_SAVE ImageResponseOrigin
}

func GetImageResponseOriginEnum() ImageResponseOriginEnum {
	return ImageResponseOriginEnum{
		CUSTOMIZE: ImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: ImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c ImageResponseOrigin) Value() string {
	return c.value
}

func (c ImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type ImageResponseResourceCategories struct {
	value string
}

type ImageResponseResourceCategoriesEnum struct {
	ASCEND ImageResponseResourceCategories
	CPU    ImageResponseResourceCategories
	GPU    ImageResponseResourceCategories
}

func GetImageResponseResourceCategoriesEnum() ImageResponseResourceCategoriesEnum {
	return ImageResponseResourceCategoriesEnum{
		ASCEND: ImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: ImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: ImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c ImageResponseResourceCategories) Value() string {
	return c.value
}

func (c ImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type ImageResponseServiceType struct {
	value string
}

type ImageResponseServiceTypeEnum struct {
	COMMON    ImageResponseServiceType
	DEV       ImageResponseServiceType
	INFERENCE ImageResponseServiceType
	TRAIN     ImageResponseServiceType
	UNKNOWN   ImageResponseServiceType
}

func GetImageResponseServiceTypeEnum() ImageResponseServiceTypeEnum {
	return ImageResponseServiceTypeEnum{
		COMMON: ImageResponseServiceType{
			value: "COMMON",
		},
		DEV: ImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: ImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: ImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: ImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c ImageResponseServiceType) Value() string {
	return c.value
}

func (c ImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type ImageResponseStatus struct {
	value string
}

type ImageResponseStatusEnum struct {
	ACTIVE        ImageResponseStatus
	CREATE_FAILED ImageResponseStatus
	CREATING      ImageResponseStatus
	DELETED       ImageResponseStatus
	ERROR         ImageResponseStatus
	INIT          ImageResponseStatus
}

func GetImageResponseStatusEnum() ImageResponseStatusEnum {
	return ImageResponseStatusEnum{
		ACTIVE: ImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: ImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: ImageResponseStatus{
			value: "CREATING",
		},
		DELETED: ImageResponseStatus{
			value: "DELETED",
		},
		ERROR: ImageResponseStatus{
			value: "ERROR",
		},
		INIT: ImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c ImageResponseStatus) Value() string {
	return c.value
}

func (c ImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type ImageResponseType struct {
	value string
}

type ImageResponseTypeEnum struct {
	BUILD_IN  ImageResponseType
	DEDICATED ImageResponseType
}

func GetImageResponseTypeEnum() ImageResponseTypeEnum {
	return ImageResponseTypeEnum{
		BUILD_IN: ImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: ImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c ImageResponseType) Value() string {
	return c.value
}

func (c ImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseType) UnmarshalJSON(b []byte) error {
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

type ImageResponseVisibility struct {
	value string
}

type ImageResponseVisibilityEnum struct {
	HIDDEN  ImageResponseVisibility
	PRIVATE ImageResponseVisibility
	PUBLIC  ImageResponseVisibility
}

func GetImageResponseVisibilityEnum() ImageResponseVisibilityEnum {
	return ImageResponseVisibilityEnum{
		HIDDEN: ImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: ImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: ImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c ImageResponseVisibility) Value() string {
	return c.value
}

func (c ImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
