package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDevServerResponse Response Object
type CreateDevServerResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：计费模式。 **取值范围**： - [COMMON：同时支持包周期和按需](tag:hws,hws_hk) - POST_PAID：按需模式 - [PRE_PAID：包周期](tag:hws,hws_hk)
	ChargingMode *CreateDevServerResponseChargingMode `json:"charging_mode,omitempty"`

	CloudServer *CloudServer `json:"cloud_server,omitempty"`

	// **参数解释**：实例私有IP信息。
	EndpointsResponse *[]Endpoints `json:"endpoints_response,omitempty"`

	// **参数解释**：实例规格名称。 **取值范围**：^.{1,128}$。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：实例ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	Id *string `json:"id,omitempty"`

	// **参数解释**：密钥对名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	KeyPairName *string `json:"key_pair_name,omitempty"`

	// **参数解释**：实例名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	Name *string `json:"name,omitempty"`

	// **参数解释**：订单ID。 **取值范围**：^[a-zA-Z0-9]{1,64}$。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：实例状态。表示实例的当前运行状态，用于监控实例的生命周期和健康状况。 **取值范围**： - CREATE_FAILED: 创建失败 - CREATING: 创建中 - DELETED: 已删除 - DELETE_FAILED: 删除失败 - DELETING: 删除中 - ERROR: 错误 - RUNNING: 运行中 - STARTING: 启动中 - START_FAILED: 启动失败 - STOPPED: 已停止 - STOPPING: 停止中 - STOP_FAILED: 停止失败 - REBOOTING: 重启中 - REBOOT_FAILED: 重启失败 - CHANGINGOS: 切换操作系统中 - CHANGINGOS_FAILED: 切换操作系统失败 - REINSTALLINGOS: 重装操作系统中 - REINSTALLINGOS_FAILED: 重装操作系统失败
	Status *CreateDevServerResponseStatus `json:"status,omitempty"`

	// **参数解释**：实例所在虚拟私有云ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：服务器私有IP信息。
	Endpoints *[]EndpointsRes `json:"endpoints,omitempty"`

	// **参数解释**：挂载硬盘信息。
	Volumes *[]ServerVolume `json:"volumes,omitempty"`

	Image *ServerImageResponse `json:"image,omitempty"`

	// **参数解释**：服务器归属类型。 **取值范围**： - [HPS：超节点服务器](tag:hws,hws_hk) - [SPOD：整柜服务器](tag:hws,hws_hk) - [SERVER：单台服务器](tag:hws,hws_hk)
	Category *CreateDevServerResponseCategory `json:"category,omitempty"`

	ServerHps *ServerHpsInfo `json:"server_hps,omitempty"`

	// **参数解释**：实例所在子网的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	SubnetId       *string `json:"subnet_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDevServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDevServerResponse struct{}"
	}

	return strings.Join([]string{"CreateDevServerResponse", string(data)}, " ")
}

type CreateDevServerResponseChargingMode struct {
	value string
}

type CreateDevServerResponseChargingModeEnum struct {
	COMMON    CreateDevServerResponseChargingMode
	POST_PAID CreateDevServerResponseChargingMode
	PRE_PAID  CreateDevServerResponseChargingMode
}

func GetCreateDevServerResponseChargingModeEnum() CreateDevServerResponseChargingModeEnum {
	return CreateDevServerResponseChargingModeEnum{
		COMMON: CreateDevServerResponseChargingMode{
			value: "COMMON",
		},
		POST_PAID: CreateDevServerResponseChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: CreateDevServerResponseChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c CreateDevServerResponseChargingMode) Value() string {
	return c.value
}

func (c CreateDevServerResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDevServerResponseChargingMode) UnmarshalJSON(b []byte) error {
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

type CreateDevServerResponseStatus struct {
	value string
}

type CreateDevServerResponseStatusEnum struct {
	CREATE_FAILED         CreateDevServerResponseStatus
	CREATING              CreateDevServerResponseStatus
	DELETED               CreateDevServerResponseStatus
	DELETE_FAILED         CreateDevServerResponseStatus
	DELETING              CreateDevServerResponseStatus
	ERROR                 CreateDevServerResponseStatus
	RUNNING               CreateDevServerResponseStatus
	STARTING              CreateDevServerResponseStatus
	START_FAILED          CreateDevServerResponseStatus
	STOPPED               CreateDevServerResponseStatus
	STOPPING              CreateDevServerResponseStatus
	STOP_FAILED           CreateDevServerResponseStatus
	REBOOTING             CreateDevServerResponseStatus
	REBOOT_FAILED         CreateDevServerResponseStatus
	CHANGINGOS            CreateDevServerResponseStatus
	CHANGINGOS_FAILED     CreateDevServerResponseStatus
	REINSTALLINGOS        CreateDevServerResponseStatus
	REINSTALLINGOS_FAILED CreateDevServerResponseStatus
}

func GetCreateDevServerResponseStatusEnum() CreateDevServerResponseStatusEnum {
	return CreateDevServerResponseStatusEnum{
		CREATE_FAILED: CreateDevServerResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: CreateDevServerResponseStatus{
			value: "CREATING",
		},
		DELETED: CreateDevServerResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: CreateDevServerResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: CreateDevServerResponseStatus{
			value: "DELETING",
		},
		ERROR: CreateDevServerResponseStatus{
			value: "ERROR",
		},
		RUNNING: CreateDevServerResponseStatus{
			value: "RUNNING",
		},
		STARTING: CreateDevServerResponseStatus{
			value: "STARTING",
		},
		START_FAILED: CreateDevServerResponseStatus{
			value: "START_FAILED",
		},
		STOPPED: CreateDevServerResponseStatus{
			value: "STOPPED",
		},
		STOPPING: CreateDevServerResponseStatus{
			value: "STOPPING",
		},
		STOP_FAILED: CreateDevServerResponseStatus{
			value: "STOP_FAILED",
		},
		REBOOTING: CreateDevServerResponseStatus{
			value: "REBOOTING",
		},
		REBOOT_FAILED: CreateDevServerResponseStatus{
			value: "REBOOT_FAILED",
		},
		CHANGINGOS: CreateDevServerResponseStatus{
			value: "CHANGINGOS",
		},
		CHANGINGOS_FAILED: CreateDevServerResponseStatus{
			value: "CHANGINGOS_FAILED",
		},
		REINSTALLINGOS: CreateDevServerResponseStatus{
			value: "REINSTALLINGOS",
		},
		REINSTALLINGOS_FAILED: CreateDevServerResponseStatus{
			value: "REINSTALLINGOS_FAILED",
		},
	}
}

func (c CreateDevServerResponseStatus) Value() string {
	return c.value
}

func (c CreateDevServerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDevServerResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateDevServerResponseCategory struct {
	value string
}

type CreateDevServerResponseCategoryEnum struct {
	SPOD   CreateDevServerResponseCategory
	SERVER CreateDevServerResponseCategory
	HPS    CreateDevServerResponseCategory
}

func GetCreateDevServerResponseCategoryEnum() CreateDevServerResponseCategoryEnum {
	return CreateDevServerResponseCategoryEnum{
		SPOD: CreateDevServerResponseCategory{
			value: "SPOD",
		},
		SERVER: CreateDevServerResponseCategory{
			value: "SERVER",
		},
		HPS: CreateDevServerResponseCategory{
			value: "HPS",
		},
	}
}

func (c CreateDevServerResponseCategory) Value() string {
	return c.value
}

func (c CreateDevServerResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDevServerResponseCategory) UnmarshalJSON(b []byte) error {
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
