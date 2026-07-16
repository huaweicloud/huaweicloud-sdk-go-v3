package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateImageResponse Response Object
type CreateImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *CreateImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]CreateImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *CreateImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]CreateImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *CreateImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *CreateImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]CreateImageResponseSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *CreateImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *CreateImageResponseVisibility `json:"visibility,omitempty"`

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
	Tags           *[]TmsTagResponse `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageResponse struct{}"
	}

	return strings.Join([]string{"CreateImageResponse", string(data)}, " ")
}

type CreateImageResponseArch struct {
	value string
}

type CreateImageResponseArchEnum struct {
	AARCH64 CreateImageResponseArch
	X86_64  CreateImageResponseArch
}

func GetCreateImageResponseArchEnum() CreateImageResponseArchEnum {
	return CreateImageResponseArchEnum{
		AARCH64: CreateImageResponseArch{
			value: "AARCH64",
		},
		X86_64: CreateImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c CreateImageResponseArch) Value() string {
	return c.value
}

func (c CreateImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseArch) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseDevServices struct {
	value string
}

type CreateImageResponseDevServicesEnum struct {
	NOTEBOOK CreateImageResponseDevServices
	SSH      CreateImageResponseDevServices
}

func GetCreateImageResponseDevServicesEnum() CreateImageResponseDevServicesEnum {
	return CreateImageResponseDevServicesEnum{
		NOTEBOOK: CreateImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: CreateImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c CreateImageResponseDevServices) Value() string {
	return c.value
}

func (c CreateImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseOrigin struct {
	value string
}

type CreateImageResponseOriginEnum struct {
	CUSTOMIZE  CreateImageResponseOrigin
	IMAGE_SAVE CreateImageResponseOrigin
}

func GetCreateImageResponseOriginEnum() CreateImageResponseOriginEnum {
	return CreateImageResponseOriginEnum{
		CUSTOMIZE: CreateImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: CreateImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c CreateImageResponseOrigin) Value() string {
	return c.value
}

func (c CreateImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseResourceCategories struct {
	value string
}

type CreateImageResponseResourceCategoriesEnum struct {
	ASCEND CreateImageResponseResourceCategories
	CPU    CreateImageResponseResourceCategories
	GPU    CreateImageResponseResourceCategories
}

func GetCreateImageResponseResourceCategoriesEnum() CreateImageResponseResourceCategoriesEnum {
	return CreateImageResponseResourceCategoriesEnum{
		ASCEND: CreateImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: CreateImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: CreateImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c CreateImageResponseResourceCategories) Value() string {
	return c.value
}

func (c CreateImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseServiceType struct {
	value string
}

type CreateImageResponseServiceTypeEnum struct {
	COMMON    CreateImageResponseServiceType
	DEV       CreateImageResponseServiceType
	INFERENCE CreateImageResponseServiceType
	TRAIN     CreateImageResponseServiceType
	UNKNOWN   CreateImageResponseServiceType
}

func GetCreateImageResponseServiceTypeEnum() CreateImageResponseServiceTypeEnum {
	return CreateImageResponseServiceTypeEnum{
		COMMON: CreateImageResponseServiceType{
			value: "COMMON",
		},
		DEV: CreateImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: CreateImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: CreateImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: CreateImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c CreateImageResponseServiceType) Value() string {
	return c.value
}

func (c CreateImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseStatus struct {
	value string
}

type CreateImageResponseStatusEnum struct {
	ACTIVE        CreateImageResponseStatus
	CREATE_FAILED CreateImageResponseStatus
	CREATING      CreateImageResponseStatus
	DELETED       CreateImageResponseStatus
	ERROR         CreateImageResponseStatus
	INIT          CreateImageResponseStatus
}

func GetCreateImageResponseStatusEnum() CreateImageResponseStatusEnum {
	return CreateImageResponseStatusEnum{
		ACTIVE: CreateImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: CreateImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: CreateImageResponseStatus{
			value: "CREATING",
		},
		DELETED: CreateImageResponseStatus{
			value: "DELETED",
		},
		ERROR: CreateImageResponseStatus{
			value: "ERROR",
		},
		INIT: CreateImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c CreateImageResponseStatus) Value() string {
	return c.value
}

func (c CreateImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseSupportResCategories struct {
	value string
}

type CreateImageResponseSupportResCategoriesEnum struct {
	ASCEND CreateImageResponseSupportResCategories
	CPU    CreateImageResponseSupportResCategories
	GPU    CreateImageResponseSupportResCategories
}

func GetCreateImageResponseSupportResCategoriesEnum() CreateImageResponseSupportResCategoriesEnum {
	return CreateImageResponseSupportResCategoriesEnum{
		ASCEND: CreateImageResponseSupportResCategories{
			value: "ASCEND",
		},
		CPU: CreateImageResponseSupportResCategories{
			value: "CPU",
		},
		GPU: CreateImageResponseSupportResCategories{
			value: "GPU",
		},
	}
}

func (c CreateImageResponseSupportResCategories) Value() string {
	return c.value
}

func (c CreateImageResponseSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseSupportResCategories) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseType struct {
	value string
}

type CreateImageResponseTypeEnum struct {
	BUILD_IN  CreateImageResponseType
	DEDICATED CreateImageResponseType
}

func GetCreateImageResponseTypeEnum() CreateImageResponseTypeEnum {
	return CreateImageResponseTypeEnum{
		BUILD_IN: CreateImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: CreateImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c CreateImageResponseType) Value() string {
	return c.value
}

func (c CreateImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseType) UnmarshalJSON(b []byte) error {
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

type CreateImageResponseVisibility struct {
	value string
}

type CreateImageResponseVisibilityEnum struct {
	HIDDEN  CreateImageResponseVisibility
	PRIVATE CreateImageResponseVisibility
	PUBLIC  CreateImageResponseVisibility
}

func GetCreateImageResponseVisibilityEnum() CreateImageResponseVisibilityEnum {
	return CreateImageResponseVisibilityEnum{
		HIDDEN: CreateImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: CreateImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: CreateImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c CreateImageResponseVisibility) Value() string {
	return c.value
}

func (c CreateImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
