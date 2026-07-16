package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteImageResponse Response Object
type DeleteImageResponse struct {

	// **参数解释**：该镜像所支持处理器架构类型。 **取值范围**：枚举类型，取值如下： - X86_64：x86处理器架构。 - AARCH64：ARM体系架构。
	Arch *DeleteImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：镜像创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：该镜像所对应的描述信息。 **取值范围**：长度限制512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像支持的服务。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：镜像支持通过https协议访问Notebook。 - SSH：镜像支持本地IDE通过SSH协议远程连接Notebook。
	DevServices *[]DeleteImageResponseDevServices `json:"dev_services,omitempty"`

	// **参数解释**：待创建Notebook实例的镜像，需要指定镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。预置镜像的ID参考[查询支持的镜像列表](ListImage.xml)获取。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名称。 **取值范围**：长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：指定镜像来源。 **取值范围**：枚举类型，取值如下： - CUSTOMIZE：用户自定义构建镜像。 - IMAGE_SAVE：Notebook实例保存镜像。
	Origin *DeleteImageResponseOrigin `json:"origin,omitempty"`

	// **参数解释**：镜像支持的规格。枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	ResourceCategories *[]DeleteImageResponseResourceCategories `json:"resource_categories,omitempty"`

	// **参数解释**：镜像支持服务类型。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE：建议仅在推理部署场景使用。 - TRAIN：建议仅在训练任务场景使用。 - DEV：建议仅在开发调测场景使用。 - UNKNOWN：未明确设置的镜像支持的服务类型。
	ServiceType *DeleteImageResponseServiceType `json:"service_type,omitempty"`

	// **参数解释**：镜像大小（单位KB）。 **取值范围**：不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：镜像状态。 **取值范围**：枚举类型，取值如下： - INIT：初始化。 - CREATING：镜像保存中，此时Notebook不可用。 - CREATE_FAILED：镜像保存失败。 - ERROR：错误。 - DELETED：已删除。 - ACTIVE：镜像保存成功，保存的镜像可以在SWR控制台查看，同时可以基于保存的镜像创建Notebook实例。
	Status *DeleteImageResponseStatus `json:"status,omitempty"`

	// **参数解释**：镜像保存操作过程中，构建信息展示。 **取值范围**：不涉及。
	StatusMessage *string `json:"status_message,omitempty"`

	// **参数解释**：镜像支持的规格。 枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	SupportResCategories *[]DeleteImageResponseSupportResCategories `json:"support_res_categories,omitempty"`

	// **参数解释**：SWR镜像地址。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：镜像Tag。 **取值范围**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：镜像类型。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。
	Type *DeleteImageResponseType `json:"type,omitempty"`

	// **参数解释**：镜像最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像可见度。 **取值范围**：枚举类型，取值如下： - PRIVATE：私有镜像。 - PUBLIC：所有用户可以根据image_id来进行只读使用。
	Visibility *DeleteImageResponseVisibility `json:"visibility,omitempty"`

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

func (o DeleteImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageResponse", string(data)}, " ")
}

type DeleteImageResponseArch struct {
	value string
}

type DeleteImageResponseArchEnum struct {
	AARCH64 DeleteImageResponseArch
	X86_64  DeleteImageResponseArch
}

func GetDeleteImageResponseArchEnum() DeleteImageResponseArchEnum {
	return DeleteImageResponseArchEnum{
		AARCH64: DeleteImageResponseArch{
			value: "AARCH64",
		},
		X86_64: DeleteImageResponseArch{
			value: "X86_64",
		},
	}
}

func (c DeleteImageResponseArch) Value() string {
	return c.value
}

func (c DeleteImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseArch) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseDevServices struct {
	value string
}

type DeleteImageResponseDevServicesEnum struct {
	NOTEBOOK DeleteImageResponseDevServices
	SSH      DeleteImageResponseDevServices
}

func GetDeleteImageResponseDevServicesEnum() DeleteImageResponseDevServicesEnum {
	return DeleteImageResponseDevServicesEnum{
		NOTEBOOK: DeleteImageResponseDevServices{
			value: "NOTEBOOK",
		},
		SSH: DeleteImageResponseDevServices{
			value: "SSH",
		},
	}
}

func (c DeleteImageResponseDevServices) Value() string {
	return c.value
}

func (c DeleteImageResponseDevServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseDevServices) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseOrigin struct {
	value string
}

type DeleteImageResponseOriginEnum struct {
	CUSTOMIZE  DeleteImageResponseOrigin
	IMAGE_SAVE DeleteImageResponseOrigin
}

func GetDeleteImageResponseOriginEnum() DeleteImageResponseOriginEnum {
	return DeleteImageResponseOriginEnum{
		CUSTOMIZE: DeleteImageResponseOrigin{
			value: "CUSTOMIZE",
		},
		IMAGE_SAVE: DeleteImageResponseOrigin{
			value: "IMAGE_SAVE",
		},
	}
}

func (c DeleteImageResponseOrigin) Value() string {
	return c.value
}

func (c DeleteImageResponseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseOrigin) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseResourceCategories struct {
	value string
}

type DeleteImageResponseResourceCategoriesEnum struct {
	ASCEND DeleteImageResponseResourceCategories
	CPU    DeleteImageResponseResourceCategories
	GPU    DeleteImageResponseResourceCategories
}

func GetDeleteImageResponseResourceCategoriesEnum() DeleteImageResponseResourceCategoriesEnum {
	return DeleteImageResponseResourceCategoriesEnum{
		ASCEND: DeleteImageResponseResourceCategories{
			value: "ASCEND",
		},
		CPU: DeleteImageResponseResourceCategories{
			value: "CPU",
		},
		GPU: DeleteImageResponseResourceCategories{
			value: "GPU",
		},
	}
}

func (c DeleteImageResponseResourceCategories) Value() string {
	return c.value
}

func (c DeleteImageResponseResourceCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseResourceCategories) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseServiceType struct {
	value string
}

type DeleteImageResponseServiceTypeEnum struct {
	COMMON    DeleteImageResponseServiceType
	DEV       DeleteImageResponseServiceType
	INFERENCE DeleteImageResponseServiceType
	TRAIN     DeleteImageResponseServiceType
	UNKNOWN   DeleteImageResponseServiceType
}

func GetDeleteImageResponseServiceTypeEnum() DeleteImageResponseServiceTypeEnum {
	return DeleteImageResponseServiceTypeEnum{
		COMMON: DeleteImageResponseServiceType{
			value: "COMMON",
		},
		DEV: DeleteImageResponseServiceType{
			value: "DEV",
		},
		INFERENCE: DeleteImageResponseServiceType{
			value: "INFERENCE",
		},
		TRAIN: DeleteImageResponseServiceType{
			value: "TRAIN",
		},
		UNKNOWN: DeleteImageResponseServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c DeleteImageResponseServiceType) Value() string {
	return c.value
}

func (c DeleteImageResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseServiceType) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseStatus struct {
	value string
}

type DeleteImageResponseStatusEnum struct {
	ACTIVE        DeleteImageResponseStatus
	CREATE_FAILED DeleteImageResponseStatus
	CREATING      DeleteImageResponseStatus
	DELETED       DeleteImageResponseStatus
	ERROR         DeleteImageResponseStatus
	INIT          DeleteImageResponseStatus
}

func GetDeleteImageResponseStatusEnum() DeleteImageResponseStatusEnum {
	return DeleteImageResponseStatusEnum{
		ACTIVE: DeleteImageResponseStatus{
			value: "ACTIVE",
		},
		CREATE_FAILED: DeleteImageResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: DeleteImageResponseStatus{
			value: "CREATING",
		},
		DELETED: DeleteImageResponseStatus{
			value: "DELETED",
		},
		ERROR: DeleteImageResponseStatus{
			value: "ERROR",
		},
		INIT: DeleteImageResponseStatus{
			value: "INIT",
		},
	}
}

func (c DeleteImageResponseStatus) Value() string {
	return c.value
}

func (c DeleteImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseSupportResCategories struct {
	value string
}

type DeleteImageResponseSupportResCategoriesEnum struct {
	ASCEND DeleteImageResponseSupportResCategories
	CPU    DeleteImageResponseSupportResCategories
	GPU    DeleteImageResponseSupportResCategories
}

func GetDeleteImageResponseSupportResCategoriesEnum() DeleteImageResponseSupportResCategoriesEnum {
	return DeleteImageResponseSupportResCategoriesEnum{
		ASCEND: DeleteImageResponseSupportResCategories{
			value: "ASCEND",
		},
		CPU: DeleteImageResponseSupportResCategories{
			value: "CPU",
		},
		GPU: DeleteImageResponseSupportResCategories{
			value: "GPU",
		},
	}
}

func (c DeleteImageResponseSupportResCategories) Value() string {
	return c.value
}

func (c DeleteImageResponseSupportResCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseSupportResCategories) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseType struct {
	value string
}

type DeleteImageResponseTypeEnum struct {
	BUILD_IN  DeleteImageResponseType
	DEDICATED DeleteImageResponseType
}

func GetDeleteImageResponseTypeEnum() DeleteImageResponseTypeEnum {
	return DeleteImageResponseTypeEnum{
		BUILD_IN: DeleteImageResponseType{
			value: "BUILD_IN",
		},
		DEDICATED: DeleteImageResponseType{
			value: "DEDICATED",
		},
	}
}

func (c DeleteImageResponseType) Value() string {
	return c.value
}

func (c DeleteImageResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseType) UnmarshalJSON(b []byte) error {
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

type DeleteImageResponseVisibility struct {
	value string
}

type DeleteImageResponseVisibilityEnum struct {
	HIDDEN  DeleteImageResponseVisibility
	PRIVATE DeleteImageResponseVisibility
	PUBLIC  DeleteImageResponseVisibility
}

func GetDeleteImageResponseVisibilityEnum() DeleteImageResponseVisibilityEnum {
	return DeleteImageResponseVisibilityEnum{
		HIDDEN: DeleteImageResponseVisibility{
			value: "HIDDEN",
		},
		PRIVATE: DeleteImageResponseVisibility{
			value: "PRIVATE",
		},
		PUBLIC: DeleteImageResponseVisibility{
			value: "PUBLIC",
		},
	}
}

func (c DeleteImageResponseVisibility) Value() string {
	return c.value
}

func (c DeleteImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageResponseVisibility) UnmarshalJSON(b []byte) error {
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
