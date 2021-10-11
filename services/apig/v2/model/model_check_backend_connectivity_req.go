package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckBackendConnectivityReq struct {
	// 后端服务配置方式 backend_address - 配置后端服务地址（不使用负载通道）  vpc_channel - 使用负载通道

	BackendType CheckBackendConnectivityReqBackendType `json:"backend_type"`
	// 后端服务地址，当type为backend_address时必填

	BackendAddress *string `json:"backend_address,omitempty"`
	// 负载通道编号，当type为vpc_channel时必填

	VpcChannelId *string `json:"vpc_channel_id,omitempty"`
}

func (o CheckBackendConnectivityReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckBackendConnectivityReq struct{}"
	}

	return strings.Join([]string{"CheckBackendConnectivityReq", string(data)}, " ")
}

type CheckBackendConnectivityReqBackendType struct {
	value string
}

type CheckBackendConnectivityReqBackendTypeEnum struct {
	BACKEND_ADDRESS CheckBackendConnectivityReqBackendType
	VPC_CHANNEL     CheckBackendConnectivityReqBackendType
}

func GetCheckBackendConnectivityReqBackendTypeEnum() CheckBackendConnectivityReqBackendTypeEnum {
	return CheckBackendConnectivityReqBackendTypeEnum{
		BACKEND_ADDRESS: CheckBackendConnectivityReqBackendType{
			value: "backend_address",
		},
		VPC_CHANNEL: CheckBackendConnectivityReqBackendType{
			value: "vpc_channel",
		},
	}
}

func (c CheckBackendConnectivityReqBackendType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CheckBackendConnectivityReqBackendType) UnmarshalJSON(b []byte) error {
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
