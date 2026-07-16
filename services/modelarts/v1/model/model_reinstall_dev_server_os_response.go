package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReinstallDevServerOsResponse Response Object
type ReinstallDevServerOsResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：计费模式。 **取值范围**： - [COMMON：同时支持包周期和按需](tag:hws,hws_hk) - POST_PAID：按需模式 - [PRE_PAID：包周期](tag:hws,hws_hk)
	ChargingMode *ReinstallDevServerOsResponseChargingMode `json:"charging_mode,omitempty"`

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
	Status *ReinstallDevServerOsResponseStatus `json:"status,omitempty"`

	// **参数解释**：实例所在虚拟私有云ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：服务器私有IP信息。
	Endpoints *[]EndpointsRes `json:"endpoints,omitempty"`

	// **参数解释**：挂载硬盘信息。
	Volumes *[]ServerVolume `json:"volumes,omitempty"`

	Image *ServerImageResponse `json:"image,omitempty"`

	// **参数解释**：服务器归属类型。 **取值范围**： - [HPS：超节点服务器](tag:hws,hws_hk) - [SPOD：整柜服务器](tag:hws,hws_hk) - [SERVER：单台服务器](tag:hws,hws_hk)
	Category *ReinstallDevServerOsResponseCategory `json:"category,omitempty"`

	ServerHps *ServerHpsInfo `json:"server_hps,omitempty"`

	// **参数解释**：实例所在子网的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	SubnetId *string `json:"subnet_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReinstallDevServerOsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallDevServerOsResponse struct{}"
	}

	return strings.Join([]string{"ReinstallDevServerOsResponse", string(data)}, " ")
}

type ReinstallDevServerOsResponseChargingMode struct {
	value string
}

type ReinstallDevServerOsResponseChargingModeEnum struct {
	COMMON    ReinstallDevServerOsResponseChargingMode
	POST_PAID ReinstallDevServerOsResponseChargingMode
	PRE_PAID  ReinstallDevServerOsResponseChargingMode
}

func GetReinstallDevServerOsResponseChargingModeEnum() ReinstallDevServerOsResponseChargingModeEnum {
	return ReinstallDevServerOsResponseChargingModeEnum{
		COMMON: ReinstallDevServerOsResponseChargingMode{
			value: "COMMON",
		},
		POST_PAID: ReinstallDevServerOsResponseChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: ReinstallDevServerOsResponseChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c ReinstallDevServerOsResponseChargingMode) Value() string {
	return c.value
}

func (c ReinstallDevServerOsResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReinstallDevServerOsResponseChargingMode) UnmarshalJSON(b []byte) error {
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

type ReinstallDevServerOsResponseStatus struct {
	value string
}

type ReinstallDevServerOsResponseStatusEnum struct {
	CREATE_FAILED         ReinstallDevServerOsResponseStatus
	CREATING              ReinstallDevServerOsResponseStatus
	DELETED               ReinstallDevServerOsResponseStatus
	DELETE_FAILED         ReinstallDevServerOsResponseStatus
	DELETING              ReinstallDevServerOsResponseStatus
	ERROR                 ReinstallDevServerOsResponseStatus
	RUNNING               ReinstallDevServerOsResponseStatus
	STARTING              ReinstallDevServerOsResponseStatus
	START_FAILED          ReinstallDevServerOsResponseStatus
	STOPPED               ReinstallDevServerOsResponseStatus
	STOPPING              ReinstallDevServerOsResponseStatus
	STOP_FAILED           ReinstallDevServerOsResponseStatus
	REBOOTING             ReinstallDevServerOsResponseStatus
	REBOOT_FAILED         ReinstallDevServerOsResponseStatus
	CHANGINGOS            ReinstallDevServerOsResponseStatus
	CHANGINGOS_FAILED     ReinstallDevServerOsResponseStatus
	REINSTALLINGOS        ReinstallDevServerOsResponseStatus
	REINSTALLINGOS_FAILED ReinstallDevServerOsResponseStatus
}

func GetReinstallDevServerOsResponseStatusEnum() ReinstallDevServerOsResponseStatusEnum {
	return ReinstallDevServerOsResponseStatusEnum{
		CREATE_FAILED: ReinstallDevServerOsResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: ReinstallDevServerOsResponseStatus{
			value: "CREATING",
		},
		DELETED: ReinstallDevServerOsResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: ReinstallDevServerOsResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: ReinstallDevServerOsResponseStatus{
			value: "DELETING",
		},
		ERROR: ReinstallDevServerOsResponseStatus{
			value: "ERROR",
		},
		RUNNING: ReinstallDevServerOsResponseStatus{
			value: "RUNNING",
		},
		STARTING: ReinstallDevServerOsResponseStatus{
			value: "STARTING",
		},
		START_FAILED: ReinstallDevServerOsResponseStatus{
			value: "START_FAILED",
		},
		STOPPED: ReinstallDevServerOsResponseStatus{
			value: "STOPPED",
		},
		STOPPING: ReinstallDevServerOsResponseStatus{
			value: "STOPPING",
		},
		STOP_FAILED: ReinstallDevServerOsResponseStatus{
			value: "STOP_FAILED",
		},
		REBOOTING: ReinstallDevServerOsResponseStatus{
			value: "REBOOTING",
		},
		REBOOT_FAILED: ReinstallDevServerOsResponseStatus{
			value: "REBOOT_FAILED",
		},
		CHANGINGOS: ReinstallDevServerOsResponseStatus{
			value: "CHANGINGOS",
		},
		CHANGINGOS_FAILED: ReinstallDevServerOsResponseStatus{
			value: "CHANGINGOS_FAILED",
		},
		REINSTALLINGOS: ReinstallDevServerOsResponseStatus{
			value: "REINSTALLINGOS",
		},
		REINSTALLINGOS_FAILED: ReinstallDevServerOsResponseStatus{
			value: "REINSTALLINGOS_FAILED",
		},
	}
}

func (c ReinstallDevServerOsResponseStatus) Value() string {
	return c.value
}

func (c ReinstallDevServerOsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReinstallDevServerOsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ReinstallDevServerOsResponseCategory struct {
	value string
}

type ReinstallDevServerOsResponseCategoryEnum struct {
	SPOD   ReinstallDevServerOsResponseCategory
	SERVER ReinstallDevServerOsResponseCategory
	HPS    ReinstallDevServerOsResponseCategory
}

func GetReinstallDevServerOsResponseCategoryEnum() ReinstallDevServerOsResponseCategoryEnum {
	return ReinstallDevServerOsResponseCategoryEnum{
		SPOD: ReinstallDevServerOsResponseCategory{
			value: "SPOD",
		},
		SERVER: ReinstallDevServerOsResponseCategory{
			value: "SERVER",
		},
		HPS: ReinstallDevServerOsResponseCategory{
			value: "HPS",
		},
	}
}

func (c ReinstallDevServerOsResponseCategory) Value() string {
	return c.value
}

func (c ReinstallDevServerOsResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReinstallDevServerOsResponseCategory) UnmarshalJSON(b []byte) error {
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
