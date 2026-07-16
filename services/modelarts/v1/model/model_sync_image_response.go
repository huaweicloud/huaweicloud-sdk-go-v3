package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncImageResponse Response Object
type SyncImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *SyncImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]SyncImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *SyncImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]SyncImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *SyncImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *SyncImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]SyncImageResponseSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *SyncImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *SyncImageResponseVisibility `json:"visibility,omitempty"`

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

func (o SyncImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncImageResponse struct{}"
	}

	return strings.Join([]string{"SyncImageResponse", string(data)}, " ")
}

type SyncImageResponseArch struct {
	value string
}

type SyncImageResponseArchEnum struct {
	AARCH64 SyncImageResponseArch
	X86_64  SyncImageResponseArch
}

func GetSyncImageResponseArchEnum() SyncImageResponseArchEnum {
	return SyncImageResponseArchEnum{
		AARCH64: SyncImageResponseArch{
			value: "AARCH64",
		},
		X86_64: SyncImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c SyncImageResponseArch) Value() string {
	return c.value
}

func (c SyncImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseArch) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseDevServices struct {
	value string
}

type SyncImageResponseDevServicesEnum struct {
	NOTEBOOK SyncImageResponseDevServices
	SSH      SyncImageResponseDevServices
}

func GetSyncImageResponseDevServicesEnum() SyncImageResponseDevServicesEnum {
	return SyncImageResponseDevServicesEnum{
		NOTEBOOK: SyncImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: SyncImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c SyncImageResponseDevServices) Value() string {
	return c.value
}

func (c SyncImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseOrigin struct {
	value string
}

type SyncImageResponseOriginEnum struct {
	CUSTOMIZE  SyncImageResponseOrigin
	IMAGE_SAVE SyncImageResponseOrigin
}

func GetSyncImageResponseOriginEnum() SyncImageResponseOriginEnum {
	return SyncImageResponseOriginEnum{
		CUSTOMIZE: SyncImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: SyncImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c SyncImageResponseOrigin) Value() string {
	return c.value
}

func (c SyncImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseResourceCategories struct {
	value string
}

type SyncImageResponseResourceCategoriesEnum struct {
	ASCEND SyncImageResponseResourceCategories
	CPU    SyncImageResponseResourceCategories
	GPU    SyncImageResponseResourceCategories
}

func GetSyncImageResponseResourceCategoriesEnum() SyncImageResponseResourceCategoriesEnum {
	return SyncImageResponseResourceCategoriesEnum{
		ASCEND: SyncImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: SyncImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: SyncImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c SyncImageResponseResourceCategories) Value() string {
	return c.value
}

func (c SyncImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseServiceType struct {
	value string
}

type SyncImageResponseServiceTypeEnum struct {
	COMMON    SyncImageResponseServiceType
	DEV       SyncImageResponseServiceType
	INFERENCE SyncImageResponseServiceType
	TRAIN     SyncImageResponseServiceType
	UNKNOWN   SyncImageResponseServiceType
}

func GetSyncImageResponseServiceTypeEnum() SyncImageResponseServiceTypeEnum {
	return SyncImageResponseServiceTypeEnum{
		COMMON: SyncImageResponseServiceType{
			value: "COMMON",
		},
		DEV: SyncImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: SyncImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: SyncImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: SyncImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c SyncImageResponseServiceType) Value() string {
	return c.value
}

func (c SyncImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseStatus struct {
	value string
}

type SyncImageResponseStatusEnum struct {
	ACTIVE        SyncImageResponseStatus
	CREATE_FAILED SyncImageResponseStatus
	CREATING      SyncImageResponseStatus
	DELETED       SyncImageResponseStatus
	ERROR         SyncImageResponseStatus
	INIT          SyncImageResponseStatus
}

func GetSyncImageResponseStatusEnum() SyncImageResponseStatusEnum {
	return SyncImageResponseStatusEnum{
		ACTIVE: SyncImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: SyncImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: SyncImageResponseStatus{
			value: "CREATING",
		},
		DELETED: SyncImageResponseStatus{
			value: "DELETED",
		},
		ERROR: SyncImageResponseStatus{
			value: "ERROR",
		},
		INIT: SyncImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c SyncImageResponseStatus) Value() string {
	return c.value
}

func (c SyncImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseSupportResCategories struct {
	value string
}

type SyncImageResponseSupportResCategoriesEnum struct {
	ASCEND SyncImageResponseSupportResCategories
	CPU    SyncImageResponseSupportResCategories
	GPU    SyncImageResponseSupportResCategories
}

func GetSyncImageResponseSupportResCategoriesEnum() SyncImageResponseSupportResCategoriesEnum {
	return SyncImageResponseSupportResCategoriesEnum{
		ASCEND: SyncImageResponseSupportResCategories{
			value: "ASCEND",
		},
		CPU: SyncImageResponseSupportResCategories{
			value: "CPU",
		},
		GPU: SyncImageResponseSupportResCategories{
			value: "GPU",
		},
	}
}

func (c SyncImageResponseSupportResCategories) Value() string {
	return c.value
}

func (c SyncImageResponseSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseSupportResCategories) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseType struct {
	value string
}

type SyncImageResponseTypeEnum struct {
	BUILD_IN  SyncImageResponseType
	DEDICATED SyncImageResponseType
}

func GetSyncImageResponseTypeEnum() SyncImageResponseTypeEnum {
	return SyncImageResponseTypeEnum{
		BUILD_IN: SyncImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: SyncImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c SyncImageResponseType) Value() string {
	return c.value
}

func (c SyncImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseType) UnmarshalJSON(b []byte) error {
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

type SyncImageResponseVisibility struct {
	value string
}

type SyncImageResponseVisibilityEnum struct {
	HIDDEN  SyncImageResponseVisibility
	PRIVATE SyncImageResponseVisibility
	PUBLIC  SyncImageResponseVisibility
}

func GetSyncImageResponseVisibilityEnum() SyncImageResponseVisibilityEnum {
	return SyncImageResponseVisibilityEnum{
		HIDDEN: SyncImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: SyncImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: SyncImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c SyncImageResponseVisibility) Value() string {
	return c.value
}

func (c SyncImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
