package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type EndpointConnection struct {

	// 连接编号
	Id string `json:"id"`

	// 连接报文标识
	MarkerId int32 `json:"marker_id"`

	// 连接创建时间。UTC时间，格式：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 连接更新时间。UTC时间，格式：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 用户的Domain ID
	DomainId string `json:"domain_id"`

	// 连接状态 - pendingAcceptance 待接受 - creating 创建中 - accepted 已接受 - rejected 已拒绝 - failed 失败 - deleting 删除中
	Status EndpointConnectionStatus `json:"status"`
}

func (o EndpointConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointConnection struct{}"
	}

	return strings.Join([]string{"EndpointConnection", string(data)}, " ")
}

type EndpointConnectionStatus struct {
	value string
}

type EndpointConnectionStatusEnum struct {
	PENDING_ACCEPTANCE EndpointConnectionStatus
	CREATING           EndpointConnectionStatus
	ACCEPTED           EndpointConnectionStatus
	REJECTED           EndpointConnectionStatus
	FAILED             EndpointConnectionStatus
	DELETING           EndpointConnectionStatus
}

func GetEndpointConnectionStatusEnum() EndpointConnectionStatusEnum {
	return EndpointConnectionStatusEnum{
		PENDING_ACCEPTANCE: EndpointConnectionStatus{
			value: "pendingAcceptance",
		},
		CREATING: EndpointConnectionStatus{
			value: "creating",
		},
		ACCEPTED: EndpointConnectionStatus{
			value: "accepted",
		},
		REJECTED: EndpointConnectionStatus{
			value: "rejected",
		},
		FAILED: EndpointConnectionStatus{
			value: "failed",
		},
		DELETING: EndpointConnectionStatus{
			value: "deleting",
		},
	}
}

func (c EndpointConnectionStatus) Value() string {
	return c.value
}

func (c EndpointConnectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointConnectionStatus) UnmarshalJSON(b []byte) error {
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
