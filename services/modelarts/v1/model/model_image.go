package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Image 实例镜像信息。
type Image struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *ImageArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]ImageDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *ImageOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]ImageResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *ImageServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *ImageStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]ImageSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *ImageType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *ImageVisibility `json:"visibility,omitempty"`

	// **参数解释**：工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：镜像的资源类型。 **取值范围**：枚举类型，取值如下： - ASCEND_SNT9：昇腾910芯片。 - ASCEND_SNT9B：昇腾910B芯片。 - ASCEND_SNT3：昇腾310芯片。
	FlavorType *string `json:"flavor_type,omitempty"`

	// 参数解释：SWR企业仓库ID。未使用SWR企业仓时该字段为null。 约束限制：不涉及。 取值范围：128位UUID。 默认取值：null。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`

	// **参数解释**：镜像展示名称，仅预置镜像具备该字段。
	ShowName *string `json:"show_name,omitempty"`

	// **参数解释**：镜像展示版本号，仅预置镜像具备该字段。
	ShowTag *string `json:"show_tag,omitempty"`

	// **参数解释**：镜像标签。
	Tags *[]TmsTagResponse `json:"tags,omitempty"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}

type ImageArch struct {
	value string
}

type ImageArchEnum struct {
	AARCH64 ImageArch
	X86_64  ImageArch
}

func GetImageArchEnum() ImageArchEnum {
	return ImageArchEnum{
		AARCH64: ImageArch{
			value: "AARCH64",
		},
		X86_64: ImageArch{
			value: "X86_64",
		},
	}
}

func (c ImageArch) Value() string {
	return c.value
}

func (c ImageArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageArch) UnmarshalJSON(b []byte) error {
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

type ImageDevServices struct {
	value string
}

type ImageDevServicesEnum struct {
	NOTEBOOK ImageDevServices
	SSH      ImageDevServices
}

func GetImageDevServicesEnum() ImageDevServicesEnum {
	return ImageDevServicesEnum{
		NOTEBOOK: ImageDevServices{
			value: "NOTEBOOK",
		},
		SSH: ImageDevServices{
			value: "SSH",
		},
	}
}

func (c ImageDevServices) Value() string {
	return c.value
}

func (c ImageDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageDevServices) UnmarshalJSON(b []byte) error {
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

type ImageOrigin struct {
	value string
}

type ImageOriginEnum struct {
	CUSTOMIZE  ImageOrigin
	IMAGE_SAVE ImageOrigin
}

func GetImageOriginEnum() ImageOriginEnum {
	return ImageOriginEnum{
		CUSTOMIZE: ImageOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: ImageOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c ImageOrigin) Value() string {
	return c.value
}

func (c ImageOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageOrigin) UnmarshalJSON(b []byte) error {
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

type ImageResourceCategories struct {
	value string
}

type ImageResourceCategoriesEnum struct {
	ASCEND ImageResourceCategories
	CPU    ImageResourceCategories
	GPU    ImageResourceCategories
}

func GetImageResourceCategoriesEnum() ImageResourceCategoriesEnum {
	return ImageResourceCategoriesEnum{
		ASCEND: ImageResourceCategories{
			value: "ASCEND",
		},
		CPU: ImageResourceCategories{
			value: "CPU",
		},
		GPU: ImageResourceCategories{
			value: "GPU",
		},
	}
}

func (c ImageResourceCategories) Value() string {
	return c.value
}

func (c ImageResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageResourceCategories) UnmarshalJSON(b []byte) error {
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

type ImageServiceType struct {
	value string
}

type ImageServiceTypeEnum struct {
	COMMON    ImageServiceType
	DEV       ImageServiceType
	INFERENCE ImageServiceType
	TRAIN     ImageServiceType
	UNKNOWN   ImageServiceType
}

func GetImageServiceTypeEnum() ImageServiceTypeEnum {
	return ImageServiceTypeEnum{
		COMMON: ImageServiceType{
			value: "COMMON",
		},
		DEV: ImageServiceType{
			value: "DEV",
		},
		INFERENCE: ImageServiceType{
			value: "INFERENCE",
		},
		TRAIN: ImageServiceType{
			value: "TRAIN",
		},
		UNKNOWN: ImageServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c ImageServiceType) Value() string {
	return c.value
}

func (c ImageServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageServiceType) UnmarshalJSON(b []byte) error {
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

type ImageStatus struct {
	value string
}

type ImageStatusEnum struct {
	ACTIVE        ImageStatus
	CREATE_FAILED ImageStatus
	CREATING      ImageStatus
	DELETED       ImageStatus
	ERROR         ImageStatus
	INIT          ImageStatus
}

func GetImageStatusEnum() ImageStatusEnum {
	return ImageStatusEnum{
		ACTIVE: ImageStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: ImageStatus{
			value: "CREATE_FAILED",
		},
		CREATING: ImageStatus{
			value: "CREATING",
		},
		DELETED: ImageStatus{
			value: "DELETED",
		},
		ERROR: ImageStatus{
			value: "ERROR",
		},
		INIT: ImageStatus{
			value: "INIT",
		},
	}
}

func (c ImageStatus) Value() string {
	return c.value
}

func (c ImageStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageStatus) UnmarshalJSON(b []byte) error {
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

type ImageSupportResCategories struct {
	value string
}

type ImageSupportResCategoriesEnum struct {
	ASCEND ImageSupportResCategories
	CPU    ImageSupportResCategories
	GPU    ImageSupportResCategories
}

func GetImageSupportResCategoriesEnum() ImageSupportResCategoriesEnum {
	return ImageSupportResCategoriesEnum{
		ASCEND: ImageSupportResCategories{
			value: "ASCEND",
		},
		CPU: ImageSupportResCategories{
			value: "CPU",
		},
		GPU: ImageSupportResCategories{
			value: "GPU",
		},
	}
}

func (c ImageSupportResCategories) Value() string {
	return c.value
}

func (c ImageSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSupportResCategories) UnmarshalJSON(b []byte) error {
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

type ImageType struct {
	value string
}

type ImageTypeEnum struct {
	BUILD_IN  ImageType
	DEDICATED ImageType
}

func GetImageTypeEnum() ImageTypeEnum {
	return ImageTypeEnum{
		BUILD_IN: ImageType{
			value: "BUILD_IN",
		},
		DEDICATED: ImageType{
			value: "DEDICATED",
		},
	}
}

func (c ImageType) Value() string {
	return c.value
}

func (c ImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageType) UnmarshalJSON(b []byte) error {
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

type ImageVisibility struct {
	value string
}

type ImageVisibilityEnum struct {
	HIDDEN  ImageVisibility
	PRIVATE ImageVisibility
	PUBLIC  ImageVisibility
}

func GetImageVisibilityEnum() ImageVisibilityEnum {
	return ImageVisibilityEnum{
		HIDDEN: ImageVisibility{
			value: "HIDDEN",
		},
		PRIVATE: ImageVisibility{
			value: "PRIVATE",
		},
		PUBLIC: ImageVisibility{
			value: "PUBLIC",
		},
	}
}

func (c ImageVisibility) Value() string {
	return c.value
}

func (c ImageVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageVisibility) UnmarshalJSON(b []byte) error {
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
