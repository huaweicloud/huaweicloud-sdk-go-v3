package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 终端节点列表
type Connection struct {
	// 终端节点的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 终端节点的报文标识。

	MarkerId *int32 `json:"marker_id,omitempty"`
	// 终端节点的创建时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	CreatedAt *string `json:"created_at,omitempty"`
	// 终端节点的更新时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	UpdatedAt *string `json:"updated_at,omitempty"`
	// 用户的Domain ID。

	DomainId *string `json:"domain_id,omitempty"`
	// 终端节点的连接状态。 ● pendingAcceptance：待接受 ● creating：创建中 ● accepted：已接受 ● rejected：已拒绝 ● failed：失败 ● deleting：删除中

	Status *ConnectionStatus `json:"status,omitempty"`
}

func (o Connection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connection struct{}"
	}

	return strings.Join([]string{"Connection", string(data)}, " ")
}

type ConnectionStatus struct {
	value string
}

type ConnectionStatusEnum struct {
	PENDING_ACCEPTANCE ConnectionStatus
	CREATING           ConnectionStatus
	ACCEPTED           ConnectionStatus
	REJECTED           ConnectionStatus
	FAILED             ConnectionStatus
	DELETING           ConnectionStatus
}

func GetConnectionStatusEnum() ConnectionStatusEnum {
	return ConnectionStatusEnum{
		PENDING_ACCEPTANCE: ConnectionStatus{
			value: "pendingAcceptance",
		},
		CREATING: ConnectionStatus{
			value: "creating",
		},
		ACCEPTED: ConnectionStatus{
			value: "accepted",
		},
		REJECTED: ConnectionStatus{
			value: "rejected",
		},
		FAILED: ConnectionStatus{
			value: "failed",
		},
		DELETING: ConnectionStatus{
			value: "deleting",
		},
	}
}

func (c ConnectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
