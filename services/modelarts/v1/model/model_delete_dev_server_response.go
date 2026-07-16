package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDevServerResponse Response Object
type DeleteDevServerResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：计费模式。 **取值范围**： - [COMMON：同时支持包周期和按需](tag:hws,hws_hk) - POST_PAID：按需模式 - [PRE_PAID：包周期](tag:hws,hws_hk)
	ChargingMode *DeleteDevServerResponseChargingMode `json:"charging_mode,omitempty"`

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
	Status *DeleteDevServerResponseStatus `json:"status,omitempty"`

	// **参数解释**：实例所在虚拟私有云ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：服务器私有IP信息。
	Endpoints *[]EndpointsRes `json:"endpoints,omitempty"`

	// **参数解释**：挂载硬盘信息。
	Volumes *[]ServerVolume `json:"volumes,omitempty"`

	Image *ServerImageResponse `json:"image,omitempty"`

	// **参数解释**：服务器归属类型。 **取值范围**： - [HPS：超节点服务器](tag:hws,hws_hk) - [SPOD：整柜服务器](tag:hws,hws_hk) - [SERVER：单台服务器](tag:hws,hws_hk)
	Category *DeleteDevServerResponseCategory `json:"category,omitempty"`

	ServerHps *ServerHpsInfo `json:"server_hps,omitempty"`

	// **参数解释**：实例所在子网的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	SubnetId       *string `json:"subnet_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDevServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDevServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteDevServerResponse", string(data)}, " ")
}

type DeleteDevServerResponseChargingMode struct {
	value string
}

type DeleteDevServerResponseChargingModeEnum struct {
	COMMON    DeleteDevServerResponseChargingMode
	POST_PAID DeleteDevServerResponseChargingMode
	PRE_PAID  DeleteDevServerResponseChargingMode
}

func GetDeleteDevServerResponseChargingModeEnum() DeleteDevServerResponseChargingModeEnum {
	return DeleteDevServerResponseChargingModeEnum{
		COMMON: DeleteDevServerResponseChargingMode{
			value: "COMMON",
		},
		POST_PAID: DeleteDevServerResponseChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: DeleteDevServerResponseChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c DeleteDevServerResponseChargingMode) Value() string {
	return c.value
}

func (c DeleteDevServerResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDevServerResponseChargingMode) UnmarshalJSON(b []byte) error {
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

type DeleteDevServerResponseStatus struct {
	value string
}

type DeleteDevServerResponseStatusEnum struct {
	CREATE_FAILED         DeleteDevServerResponseStatus
	CREATING              DeleteDevServerResponseStatus
	DELETED               DeleteDevServerResponseStatus
	DELETE_FAILED         DeleteDevServerResponseStatus
	DELETING              DeleteDevServerResponseStatus
	ERROR                 DeleteDevServerResponseStatus
	RUNNING               DeleteDevServerResponseStatus
	STARTING              DeleteDevServerResponseStatus
	START_FAILED          DeleteDevServerResponseStatus
	STOPPED               DeleteDevServerResponseStatus
	STOPPING              DeleteDevServerResponseStatus
	STOP_FAILED           DeleteDevServerResponseStatus
	REBOOTING             DeleteDevServerResponseStatus
	REBOOT_FAILED         DeleteDevServerResponseStatus
	CHANGINGOS            DeleteDevServerResponseStatus
	CHANGINGOS_FAILED     DeleteDevServerResponseStatus
	REINSTALLINGOS        DeleteDevServerResponseStatus
	REINSTALLINGOS_FAILED DeleteDevServerResponseStatus
}

func GetDeleteDevServerResponseStatusEnum() DeleteDevServerResponseStatusEnum {
	return DeleteDevServerResponseStatusEnum{
		CREATE_FAILED: DeleteDevServerResponseStatus{
			value: "CREATE_FAILED",
		},
		CREATING: DeleteDevServerResponseStatus{
			value: "CREATING",
		},
		DELETED: DeleteDevServerResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: DeleteDevServerResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: DeleteDevServerResponseStatus{
			value: "DELETING",
		},
		ERROR: DeleteDevServerResponseStatus{
			value: "ERROR",
		},
		RUNNING: DeleteDevServerResponseStatus{
			value: "RUNNING",
		},
		STARTING: DeleteDevServerResponseStatus{
			value: "STARTING",
		},
		START_FAILED: DeleteDevServerResponseStatus{
			value: "START_FAILED",
		},
		STOPPED: DeleteDevServerResponseStatus{
			value: "STOPPED",
		},
		STOPPING: DeleteDevServerResponseStatus{
			value: "STOPPING",
		},
		STOP_FAILED: DeleteDevServerResponseStatus{
			value: "STOP_FAILED",
		},
		REBOOTING: DeleteDevServerResponseStatus{
			value: "REBOOTING",
		},
		REBOOT_FAILED: DeleteDevServerResponseStatus{
			value: "REBOOT_FAILED",
		},
		CHANGINGOS: DeleteDevServerResponseStatus{
			value: "CHANGINGOS",
		},
		CHANGINGOS_FAILED: DeleteDevServerResponseStatus{
			value: "CHANGINGOS_FAILED",
		},
		REINSTALLINGOS: DeleteDevServerResponseStatus{
			value: "REINSTALLINGOS",
		},
		REINSTALLINGOS_FAILED: DeleteDevServerResponseStatus{
			value: "REINSTALLINGOS_FAILED",
		},
	}
}

func (c DeleteDevServerResponseStatus) Value() string {
	return c.value
}

func (c DeleteDevServerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDevServerResponseStatus) UnmarshalJSON(b []byte) error {
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

type DeleteDevServerResponseCategory struct {
	value string
}

type DeleteDevServerResponseCategoryEnum struct {
	SPOD   DeleteDevServerResponseCategory
	SERVER DeleteDevServerResponseCategory
	HPS    DeleteDevServerResponseCategory
}

func GetDeleteDevServerResponseCategoryEnum() DeleteDevServerResponseCategoryEnum {
	return DeleteDevServerResponseCategoryEnum{
		SPOD: DeleteDevServerResponseCategory{
			value: "SPOD",
		},
		SERVER: DeleteDevServerResponseCategory{
			value: "SERVER",
		},
		HPS: DeleteDevServerResponseCategory{
			value: "HPS",
		},
	}
}

func (c DeleteDevServerResponseCategory) Value() string {
	return c.value
}

func (c DeleteDevServerResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDevServerResponseCategory) UnmarshalJSON(b []byte) error {
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
