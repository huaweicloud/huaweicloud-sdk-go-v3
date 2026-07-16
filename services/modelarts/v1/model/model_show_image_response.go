package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowImageResponse Response Object
type ShowImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *ShowImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]ShowImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *ShowImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]ShowImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *ShowImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *ShowImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]ShowImageResponseSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *ShowImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *ShowImageResponseVisibility `json:"visibility,omitempty"`

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

func (o ShowImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageResponse struct{}"
	}

	return strings.Join([]string{"ShowImageResponse", string(data)}, " ")
}

type ShowImageResponseArch struct {
	value string
}

type ShowImageResponseArchEnum struct {
	AARCH64 ShowImageResponseArch
	X86_64  ShowImageResponseArch
}

func GetShowImageResponseArchEnum() ShowImageResponseArchEnum {
	return ShowImageResponseArchEnum{
		AARCH64: ShowImageResponseArch{
			value: "AARCH64",
		},
		X86_64: ShowImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c ShowImageResponseArch) Value() string {
	return c.value
}

func (c ShowImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseArch) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseDevServices struct {
	value string
}

type ShowImageResponseDevServicesEnum struct {
	NOTEBOOK ShowImageResponseDevServices
	SSH      ShowImageResponseDevServices
}

func GetShowImageResponseDevServicesEnum() ShowImageResponseDevServicesEnum {
	return ShowImageResponseDevServicesEnum{
		NOTEBOOK: ShowImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: ShowImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c ShowImageResponseDevServices) Value() string {
	return c.value
}

func (c ShowImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseOrigin struct {
	value string
}

type ShowImageResponseOriginEnum struct {
	CUSTOMIZE  ShowImageResponseOrigin
	IMAGE_SAVE ShowImageResponseOrigin
}

func GetShowImageResponseOriginEnum() ShowImageResponseOriginEnum {
	return ShowImageResponseOriginEnum{
		CUSTOMIZE: ShowImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: ShowImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c ShowImageResponseOrigin) Value() string {
	return c.value
}

func (c ShowImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseResourceCategories struct {
	value string
}

type ShowImageResponseResourceCategoriesEnum struct {
	ASCEND ShowImageResponseResourceCategories
	CPU    ShowImageResponseResourceCategories
	GPU    ShowImageResponseResourceCategories
}

func GetShowImageResponseResourceCategoriesEnum() ShowImageResponseResourceCategoriesEnum {
	return ShowImageResponseResourceCategoriesEnum{
		ASCEND: ShowImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: ShowImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: ShowImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c ShowImageResponseResourceCategories) Value() string {
	return c.value
}

func (c ShowImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseServiceType struct {
	value string
}

type ShowImageResponseServiceTypeEnum struct {
	COMMON    ShowImageResponseServiceType
	DEV       ShowImageResponseServiceType
	INFERENCE ShowImageResponseServiceType
	TRAIN     ShowImageResponseServiceType
	UNKNOWN   ShowImageResponseServiceType
}

func GetShowImageResponseServiceTypeEnum() ShowImageResponseServiceTypeEnum {
	return ShowImageResponseServiceTypeEnum{
		COMMON: ShowImageResponseServiceType{
			value: "COMMON",
		},
		DEV: ShowImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: ShowImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: ShowImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: ShowImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c ShowImageResponseServiceType) Value() string {
	return c.value
}

func (c ShowImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseStatus struct {
	value string
}

type ShowImageResponseStatusEnum struct {
	ACTIVE        ShowImageResponseStatus
	CREATE_FAILED ShowImageResponseStatus
	CREATING      ShowImageResponseStatus
	DELETED       ShowImageResponseStatus
	ERROR         ShowImageResponseStatus
	INIT          ShowImageResponseStatus
}

func GetShowImageResponseStatusEnum() ShowImageResponseStatusEnum {
	return ShowImageResponseStatusEnum{
		ACTIVE: ShowImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: ShowImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: ShowImageResponseStatus{
			value: "CREATING",
		},
		DELETED: ShowImageResponseStatus{
			value: "DELETED",
		},
		ERROR: ShowImageResponseStatus{
			value: "ERROR",
		},
		INIT: ShowImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowImageResponseStatus) Value() string {
	return c.value
}

func (c ShowImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseSupportResCategories struct {
	value string
}

type ShowImageResponseSupportResCategoriesEnum struct {
	ASCEND ShowImageResponseSupportResCategories
	CPU    ShowImageResponseSupportResCategories
	GPU    ShowImageResponseSupportResCategories
}

func GetShowImageResponseSupportResCategoriesEnum() ShowImageResponseSupportResCategoriesEnum {
	return ShowImageResponseSupportResCategoriesEnum{
		ASCEND: ShowImageResponseSupportResCategories{
			value: "ASCEND",
		},
		CPU: ShowImageResponseSupportResCategories{
			value: "CPU",
		},
		GPU: ShowImageResponseSupportResCategories{
			value: "GPU",
		},
	}
}

func (c ShowImageResponseSupportResCategories) Value() string {
	return c.value
}

func (c ShowImageResponseSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseSupportResCategories) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseType struct {
	value string
}

type ShowImageResponseTypeEnum struct {
	BUILD_IN  ShowImageResponseType
	DEDICATED ShowImageResponseType
}

func GetShowImageResponseTypeEnum() ShowImageResponseTypeEnum {
	return ShowImageResponseTypeEnum{
		BUILD_IN: ShowImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: ShowImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c ShowImageResponseType) Value() string {
	return c.value
}

func (c ShowImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseType) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseVisibility struct {
	value string
}

type ShowImageResponseVisibilityEnum struct {
	HIDDEN  ShowImageResponseVisibility
	PRIVATE ShowImageResponseVisibility
	PUBLIC  ShowImageResponseVisibility
}

func GetShowImageResponseVisibilityEnum() ShowImageResponseVisibilityEnum {
	return ShowImageResponseVisibilityEnum{
		HIDDEN: ShowImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: ShowImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: ShowImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c ShowImageResponseVisibility) Value() string {
	return c.value
}

func (c ShowImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
