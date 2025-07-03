package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NovaListServersRequest Request Object
type NovaListServersRequest struct {

	// 云服务器上次更新状态的时间戳信息。格式符合ISO 8601，例如：  CCYY-MM-DDThh:mm:ss+/-hh:mm  > 说明： >  > 可以查询到“deleted”状态的弹性云服务器。
	ChangesSince *string `json:"changes-since,omitempty"`

	// 云服务器类型ID。
	Flavor *string `json:"flavor,omitempty"`

	// 主机节点名称。
	Host *string `json:"host,omitempty"`

	// 镜像ID。  在使用image作为条件过滤时，不能同时支持其他过滤条件和分页条件。如果同时指定image及其他条件，则以image条件为准；当条件不含image时，接口功能不受限制。
	Image *string `json:"image,omitempty"`

	// IPv4地址过滤结果，匹配规则为模糊匹配。
	Ip *string `json:"ip,omitempty"`

	// 查询返回云服务器数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 从marker指定的云服务器ID的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 云服务器名称。
	Name *string `json:"name,omitempty"`

	// 查询tag字段中不包含该值的云服务器，值为标签的Key。  > 说明： >  > 系统近期对标签功能进行了升级。如果之前添加的Tag为“Key.Value”的形式，则查询的时候需要使用“Key”来查询。 >  > 例如：之前添加的tag为“a.b”,则升级后，查询时需使用“not-tags=a”。
	NotTags *string `json:"not-tags,omitempty"`

	// 批量创建弹性云服务器时，指定返回的ID，用于查询本次批量创建的弹性云服务器。
	ReservationId *string `json:"reservation_id,omitempty"`

	// 查询结果按弹性云服务器属性排序，默认排序顺序为created_at逆序。
	SortKey *NovaListServersRequestSortKey `json:"sort_key,omitempty"`

	// 云服务器状态。  取值范围： ACTIVE， BUILD，DELETED，ERROR，HARD_REBOOT，MIGRATING，REBOOT，RESIZE，REVERT_RESIZE，SHELVED，SHELVED_OFFLOADED，SHUTOFF，UNKNOWN，VERIFY_RESIZE 其中DELETED 状态虚拟机只有管理员有权限查 直到2.37微版本，非上面范围的status字段将返回空列表，2.38之后的微版本，将返回400错误。  云服务器状态说明请参考[云服务器状态](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)。
	Status *NovaListServersRequestStatus `json:"status,omitempty"`

	// 云服务器标签。
	Tags *string `json:"tags,omitempty"`

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServersRequest struct{}"
	}

	return strings.Join([]string{"NovaListServersRequest", string(data)}, " ")
}

type NovaListServersRequestSortKey struct {
	value string
}

type NovaListServersRequestSortKeyEnum struct {
	CREATED_AT        NovaListServersRequestSortKey
	AVAILABILITY_ZONE NovaListServersRequestSortKey
	DISPLAY_NAME      NovaListServersRequestSortKey
	HOST              NovaListServersRequestSortKey
	INSTANCE_TYPE_ID  NovaListServersRequestSortKey
	KEY_NAME          NovaListServersRequestSortKey
	PROJECT_ID        NovaListServersRequestSortKey
	USER_ID           NovaListServersRequestSortKey
	UPDATED_AT        NovaListServersRequestSortKey
	UUID              NovaListServersRequestSortKey
	VM_STATE          NovaListServersRequestSortKey
}

func GetNovaListServersRequestSortKeyEnum() NovaListServersRequestSortKeyEnum {
	return NovaListServersRequestSortKeyEnum{
		CREATED_AT: NovaListServersRequestSortKey{
			value: "created_at",
		},
		AVAILABILITY_ZONE: NovaListServersRequestSortKey{
			value: "availability_zone",
		},
		DISPLAY_NAME: NovaListServersRequestSortKey{
			value: "display_name",
		},
		HOST: NovaListServersRequestSortKey{
			value: "host",
		},
		INSTANCE_TYPE_ID: NovaListServersRequestSortKey{
			value: "instance_type_id",
		},
		KEY_NAME: NovaListServersRequestSortKey{
			value: "key_name",
		},
		PROJECT_ID: NovaListServersRequestSortKey{
			value: "project_id",
		},
		USER_ID: NovaListServersRequestSortKey{
			value: "user_id",
		},
		UPDATED_AT: NovaListServersRequestSortKey{
			value: "updated_at",
		},
		UUID: NovaListServersRequestSortKey{
			value: "uuid",
		},
		VM_STATE: NovaListServersRequestSortKey{
			value: "vm_state",
		},
	}
}

func (c NovaListServersRequestSortKey) Value() string {
	return c.value
}

func (c NovaListServersRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaListServersRequestSortKey) UnmarshalJSON(b []byte) error {
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

type NovaListServersRequestStatus struct {
	value string
}

type NovaListServersRequestStatusEnum struct {
	ACTIVE            NovaListServersRequestStatus
	BUILD             NovaListServersRequestStatus
	DELETED           NovaListServersRequestStatus
	ERROR             NovaListServersRequestStatus
	HARD_REBOOT       NovaListServersRequestStatus
	MIGRATING         NovaListServersRequestStatus
	REBOOT            NovaListServersRequestStatus
	RESIZE            NovaListServersRequestStatus
	REVERT_RESIZE     NovaListServersRequestStatus
	SHELVED           NovaListServersRequestStatus
	SHELVED_OFFLOADED NovaListServersRequestStatus
	SHUTOFF           NovaListServersRequestStatus
	UNKNOWN           NovaListServersRequestStatus
	VERIFY_RESIZE     NovaListServersRequestStatus
}

func GetNovaListServersRequestStatusEnum() NovaListServersRequestStatusEnum {
	return NovaListServersRequestStatusEnum{
		ACTIVE: NovaListServersRequestStatus{
			value: "ACTIVE",
		},
		BUILD: NovaListServersRequestStatus{
			value: "BUILD",
		},
		DELETED: NovaListServersRequestStatus{
			value: "DELETED",
		},
		ERROR: NovaListServersRequestStatus{
			value: "ERROR",
		},
		HARD_REBOOT: NovaListServersRequestStatus{
			value: "HARD_REBOOT",
		},
		MIGRATING: NovaListServersRequestStatus{
			value: "MIGRATING",
		},
		REBOOT: NovaListServersRequestStatus{
			value: "REBOOT",
		},
		RESIZE: NovaListServersRequestStatus{
			value: "RESIZE",
		},
		REVERT_RESIZE: NovaListServersRequestStatus{
			value: "REVERT_RESIZE",
		},
		SHELVED: NovaListServersRequestStatus{
			value: "SHELVED",
		},
		SHELVED_OFFLOADED: NovaListServersRequestStatus{
			value: "SHELVED_OFFLOADED",
		},
		SHUTOFF: NovaListServersRequestStatus{
			value: "SHUTOFF",
		},
		UNKNOWN: NovaListServersRequestStatus{
			value: "UNKNOWN",
		},
		VERIFY_RESIZE: NovaListServersRequestStatus{
			value: "VERIFY_RESIZE",
		},
	}
}

func (c NovaListServersRequestStatus) Value() string {
	return c.value
}

func (c NovaListServersRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaListServersRequestStatus) UnmarshalJSON(b []byte) error {
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
