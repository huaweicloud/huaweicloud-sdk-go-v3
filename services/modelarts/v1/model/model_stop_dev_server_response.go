package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopDevServerResponse Response Object
type StopDevServerResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：计费模式。 **取值范围**： - [COMMON：同时支持包周期和按需](tag:hws,hws_hk) - POST_PAID：按需模式 - [PRE_PAID：包周期](tag:hws,hws_hk)
	ChargingMode *StopDevServerResponseChargingMode `json:"charging_mode,omitempty"`

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
	Status *StopDevServerResponseStatus `json:"status,omitempty"`

	// **参数解释**：实例所在虚拟私有云ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：服务器私有IP信息。
	Endpoints *[]EndpointsRes `json:"endpoints,omitempty"`

	// **参数解释**：挂载硬盘信息。
	Volumes *[]ServerVolume `json:"volumes,omitempty"`

	Image *ServerImageResponse `json:"image,omitempty"`

	// **参数解释**：服务器归属类型。 **取值范围**： - [HPS：超节点服务器](tag:hws,hws_hk) - [SPOD：整柜服务器](tag:hws,hws_hk) - [SERVER：单台服务器](tag:hws,hws_hk)
	Category *StopDevServerResponseCategory `json:"category,omitempty"`

	ServerHps *ServerHpsInfo `json:"server_hps,omitempty"`

	// **参数解释**：实例所在子网的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	SubnetId       *string `json:"subnet_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopDevServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDevServerResponse struct{}"
	}

	return strings.Join([]string{"StopDevServerResponse", string(data)}, " ")
}

type StopDevServerResponseChargingMode struct {
	value string
}

type StopDevServerResponseChargingModeEnum struct {
	COMMON    StopDevServerResponseChargingMode
	POST_PAID StopDevServerResponseChargingMode
	PRE_PAID  StopDevServerResponseChargingMode
}

func GetStopDevServerResponseChargingModeEnum() StopDevServerResponseChargingModeEnum {
	return StopDevServerResponseChargingModeEnum{
		COMMON: StopDevServerResponseChargingMode{
			value: "COMMON",
		},
		POST_PAID: StopDevServerResponseChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: StopDevServerResponseChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c StopDevServerResponseChargingMode) Value() string {
	return c.value
}

func (c StopDevServerResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopDevServerResponseChargingMode) UnmarshalJSON(b []byte) error {
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

type StopDevServerResponseStatus struct {
	value string
}

type StopDevServerResponseStatusEnum struct {
	CREATE_FAILED         StopDevServerResponseStatus
	CREATING              StopDevServerResponseStatus
	DELETED               StopDevServerResponseStatus
	DELETE_FAILED         StopDevServerResponseStatus
	DELETING              StopDevServerResponseStatus
	ERROR                 StopDevServerResponseStatus
	RUNNING               StopDevServerResponseStatus
	STARTING              StopDevServerResponseStatus
	START_FAILED          StopDevServerResponseStatus
	STOPPED               StopDevServerResponseStatus
	STOPPING              StopDevServerResponseStatus
	STOP_FAILED           StopDevServerResponseStatus
	REBOOTING             StopDevServerResponseStatus
	REBOOT_FAILED         StopDevServerResponseStatus
	CHANGINGOS            StopDevServerResponseStatus
	CHANGINGOS_FAILED     StopDevServerResponseStatus
	REINSTALLINGOS        StopDevServerResponseStatus
	REINSTALLINGOS_FAILED StopDevServerResponseStatus
}

func GetStopDevServerResponseStatusEnum() StopDevServerResponseStatusEnum {
	return StopDevServerResponseStatusEnum{
		CREATE_FAILED: StopDevServerResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: StopDevServerResponseStatus{
			value: "CREATING",
		},
		DELETED: StopDevServerResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: StopDevServerResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: StopDevServerResponseStatus{
			value: "DELETING",
		},
		ERROR: StopDevServerResponseStatus{
			value: "ERROR",
		},
		RUNNING: StopDevServerResponseStatus{
			value: "RUNNING",
		},
		STARTING: StopDevServerResponseStatus{
			value: "STARTING",
		},
		START_FAILED: StopDevServerResponseStatus{
			value: "START_FAILED",
		},
		STOPPED: StopDevServerResponseStatus{
			value: "STOPPED",
		},
		STOPPING: StopDevServerResponseStatus{
			value: "STOPPING",
		},
		STOP_FAILED: StopDevServerResponseStatus{
			value: "STOP_FAILED",
		},
		REBOOTING: StopDevServerResponseStatus{
			value: "REBOOTING",
		},
		REBOOT_FAILED: StopDevServerResponseStatus{
			value: "REBOOT_FAILED",
		},
		CHANGINGOS: StopDevServerResponseStatus{
			value: "CHANGINGOS",
		},
		CHANGINGOS_FAILED: StopDevServerResponseStatus{
			value: "CHANGINGOS_FAILED",
		},
		REINSTALLINGOS: StopDevServerResponseStatus{
			value: "REINSTALLINGOS",
		},
		REINSTALLINGOS_FAILED: StopDevServerResponseStatus{
			value: "REINSTALLINGOS_FAILED",
		},
	}
}

func (c StopDevServerResponseStatus) Value() string {
	return c.value
}

func (c StopDevServerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopDevServerResponseStatus) UnmarshalJSON(b []byte) error {
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

type StopDevServerResponseCategory struct {
	value string
}

type StopDevServerResponseCategoryEnum struct {
	SPOD   StopDevServerResponseCategory
	SERVER StopDevServerResponseCategory
	HPS    StopDevServerResponseCategory
}

func GetStopDevServerResponseCategoryEnum() StopDevServerResponseCategoryEnum {
	return StopDevServerResponseCategoryEnum{
		SPOD: StopDevServerResponseCategory{
			value: "SPOD",
		},
		SERVER: StopDevServerResponseCategory{
			value: "SERVER",
		},
		HPS: StopDevServerResponseCategory{
			value: "HPS",
		},
	}
}

func (c StopDevServerResponseCategory) Value() string {
	return c.value
}

func (c StopDevServerResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopDevServerResponseCategory) UnmarshalJSON(b []byte) error {
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
