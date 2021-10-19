package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// 终端节点列表
type Endpoints struct {
	// 终端节点的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 终端节点的报文标识。

	MarkerId *int32 `json:"marker_id,omitempty"`
	// 终端节点的创建时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH: MM:SSZ

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 终端节点的更新时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH: MM:SSZ

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
	// 用户的Domain ID。

	DomainId *string `json:"domain_id,omitempty"`
	// 终端节点的连接状态。 ● pendingAcceptance：待接受 ● creating：创建中 ● accepted：已接受 ● rejected：已拒绝 ● failed：失败 ● deleting：删除中

	Status *EndpointsStatus `json:"status,omitempty"`

	Error *QueryError `json:"error,omitempty"`
}

func (o Endpoints) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Endpoints struct{}"
	}

	return strings.Join([]string{"Endpoints", string(data)}, " ")
}

type EndpointsStatus struct {
	value string
}

type EndpointsStatusEnum struct {
	PENDING_ACCEPTANCE EndpointsStatus
	CREATING           EndpointsStatus
	ACCEPTED           EndpointsStatus
	REJECTED           EndpointsStatus
	FAILED             EndpointsStatus
	DELETING           EndpointsStatus
}

func GetEndpointsStatusEnum() EndpointsStatusEnum {
	return EndpointsStatusEnum{
		PENDING_ACCEPTANCE: EndpointsStatus{
			value: "pendingAcceptance",
		},
		CREATING: EndpointsStatus{
			value: "creating",
		},
		ACCEPTED: EndpointsStatus{
			value: "accepted",
		},
		REJECTED: EndpointsStatus{
			value: "rejected",
		},
		FAILED: EndpointsStatus{
			value: "failed",
		},
		DELETING: EndpointsStatus{
			value: "deleting",
		},
	}
}

func (c EndpointsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EndpointsStatus) UnmarshalJSON(b []byte) error {
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
