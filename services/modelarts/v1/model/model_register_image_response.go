package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RegisterImageResponse Response Object
type RegisterImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *RegisterImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]RegisterImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *RegisterImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]RegisterImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *RegisterImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *RegisterImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]RegisterImageResponseSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *RegisterImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *RegisterImageResponseVisibility `json:"visibility,omitempty"`

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

func (o RegisterImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageResponse struct{}"
	}

	return strings.Join([]string{"RegisterImageResponse", string(data)}, " ")
}

type RegisterImageResponseArch struct {
	value string
}

type RegisterImageResponseArchEnum struct {
	AARCH64 RegisterImageResponseArch
	X86_64  RegisterImageResponseArch
}

func GetRegisterImageResponseArchEnum() RegisterImageResponseArchEnum {
	return RegisterImageResponseArchEnum{
		AARCH64: RegisterImageResponseArch{
			value: "AARCH64",
		},
		X86_64: RegisterImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c RegisterImageResponseArch) Value() string {
	return c.value
}

func (c RegisterImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseArch) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseDevServices struct {
	value string
}

type RegisterImageResponseDevServicesEnum struct {
	NOTEBOOK RegisterImageResponseDevServices
	SSH      RegisterImageResponseDevServices
}

func GetRegisterImageResponseDevServicesEnum() RegisterImageResponseDevServicesEnum {
	return RegisterImageResponseDevServicesEnum{
		NOTEBOOK: RegisterImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: RegisterImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c RegisterImageResponseDevServices) Value() string {
	return c.value
}

func (c RegisterImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseOrigin struct {
	value string
}

type RegisterImageResponseOriginEnum struct {
	CUSTOMIZE  RegisterImageResponseOrigin
	IMAGE_SAVE RegisterImageResponseOrigin
}

func GetRegisterImageResponseOriginEnum() RegisterImageResponseOriginEnum {
	return RegisterImageResponseOriginEnum{
		CUSTOMIZE: RegisterImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: RegisterImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c RegisterImageResponseOrigin) Value() string {
	return c.value
}

func (c RegisterImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseResourceCategories struct {
	value string
}

type RegisterImageResponseResourceCategoriesEnum struct {
	ASCEND RegisterImageResponseResourceCategories
	CPU    RegisterImageResponseResourceCategories
	GPU    RegisterImageResponseResourceCategories
}

func GetRegisterImageResponseResourceCategoriesEnum() RegisterImageResponseResourceCategoriesEnum {
	return RegisterImageResponseResourceCategoriesEnum{
		ASCEND: RegisterImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: RegisterImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: RegisterImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c RegisterImageResponseResourceCategories) Value() string {
	return c.value
}

func (c RegisterImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseServiceType struct {
	value string
}

type RegisterImageResponseServiceTypeEnum struct {
	COMMON    RegisterImageResponseServiceType
	DEV       RegisterImageResponseServiceType
	INFERENCE RegisterImageResponseServiceType
	TRAIN     RegisterImageResponseServiceType
	UNKNOWN   RegisterImageResponseServiceType
}

func GetRegisterImageResponseServiceTypeEnum() RegisterImageResponseServiceTypeEnum {
	return RegisterImageResponseServiceTypeEnum{
		COMMON: RegisterImageResponseServiceType{
			value: "COMMON",
		},
		DEV: RegisterImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: RegisterImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: RegisterImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: RegisterImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c RegisterImageResponseServiceType) Value() string {
	return c.value
}

func (c RegisterImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseStatus struct {
	value string
}

type RegisterImageResponseStatusEnum struct {
	ACTIVE        RegisterImageResponseStatus
	CREATE_FAILED RegisterImageResponseStatus
	CREATING      RegisterImageResponseStatus
	DELETED       RegisterImageResponseStatus
	ERROR         RegisterImageResponseStatus
	INIT          RegisterImageResponseStatus
}

func GetRegisterImageResponseStatusEnum() RegisterImageResponseStatusEnum {
	return RegisterImageResponseStatusEnum{
		ACTIVE: RegisterImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: RegisterImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: RegisterImageResponseStatus{
			value: "CREATING",
		},
		DELETED: RegisterImageResponseStatus{
			value: "DELETED",
		},
		ERROR: RegisterImageResponseStatus{
			value: "ERROR",
		},
		INIT: RegisterImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c RegisterImageResponseStatus) Value() string {
	return c.value
}

func (c RegisterImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseSupportResCategories struct {
	value string
}

type RegisterImageResponseSupportResCategoriesEnum struct {
	ASCEND RegisterImageResponseSupportResCategories
	CPU    RegisterImageResponseSupportResCategories
	GPU    RegisterImageResponseSupportResCategories
}

func GetRegisterImageResponseSupportResCategoriesEnum() RegisterImageResponseSupportResCategoriesEnum {
	return RegisterImageResponseSupportResCategoriesEnum{
		ASCEND: RegisterImageResponseSupportResCategories{
			value: "ASCEND",
		},
		CPU: RegisterImageResponseSupportResCategories{
			value: "CPU",
		},
		GPU: RegisterImageResponseSupportResCategories{
			value: "GPU",
		},
	}
}

func (c RegisterImageResponseSupportResCategories) Value() string {
	return c.value
}

func (c RegisterImageResponseSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseSupportResCategories) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseType struct {
	value string
}

type RegisterImageResponseTypeEnum struct {
	BUILD_IN  RegisterImageResponseType
	DEDICATED RegisterImageResponseType
}

func GetRegisterImageResponseTypeEnum() RegisterImageResponseTypeEnum {
	return RegisterImageResponseTypeEnum{
		BUILD_IN: RegisterImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: RegisterImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c RegisterImageResponseType) Value() string {
	return c.value
}

func (c RegisterImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseType) UnmarshalJSON(b []byte) error {
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

type RegisterImageResponseVisibility struct {
	value string
}

type RegisterImageResponseVisibilityEnum struct {
	HIDDEN  RegisterImageResponseVisibility
	PRIVATE RegisterImageResponseVisibility
	PUBLIC  RegisterImageResponseVisibility
}

func GetRegisterImageResponseVisibilityEnum() RegisterImageResponseVisibilityEnum {
	return RegisterImageResponseVisibilityEnum{
		HIDDEN: RegisterImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: RegisterImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: RegisterImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c RegisterImageResponseVisibility) Value() string {
	return c.value
}

func (c RegisterImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
