package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionItem struct {

	// VPC终端节点ID
	Id *string `json:"id,omitempty"`

	// VPC终端节点所属的租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// VPC终端节点所属的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// VPC终端节点的连接状态 取值范围: - pendingAcceptance:待接受 - creating:创建中 - accepted:已接受 - rejected:已拒绝 - failed:失败 - deleting:删除中
	Status *ConnectionItemStatus `json:"status,omitempty"`

	// VPC终端节点的创建时间。采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// VPC终端节点的更新时间。采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否为保护内网访问连接；如果为true则不允许添加或者移除
	Protected *bool `json:"protected,omitempty"`
}

func (o ConnectionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionItem struct{}"
	}

	return strings.Join([]string{"ConnectionItem", string(data)}, " ")
}

type ConnectionItemStatus struct {
	value string
}

type ConnectionItemStatusEnum struct {
	PENDING_ACCEPTANCE ConnectionItemStatus
	CREATING           ConnectionItemStatus
	ACCEPTED           ConnectionItemStatus
	REJECTED           ConnectionItemStatus
	FAILED             ConnectionItemStatus
	DELETING           ConnectionItemStatus
}

func GetConnectionItemStatusEnum() ConnectionItemStatusEnum {
	return ConnectionItemStatusEnum{
		PENDING_ACCEPTANCE: ConnectionItemStatus{
			value: "pendingAcceptance",
		},
		CREATING: ConnectionItemStatus{
			value: "creating",
		},
		ACCEPTED: ConnectionItemStatus{
			value: "accepted",
		},
		REJECTED: ConnectionItemStatus{
			value: "rejected",
		},
		FAILED: ConnectionItemStatus{
			value: "failed",
		},
		DELETING: ConnectionItemStatus{
			value: "deleting",
		},
	}
}

func (c ConnectionItemStatus) Value() string {
	return c.value
}

func (c ConnectionItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionItemStatus) UnmarshalJSON(b []byte) error {
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
