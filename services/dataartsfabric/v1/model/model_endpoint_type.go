package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointType Endpoint的类型  - service：Service EP，代表一个可接收Service请求资源组  - ray：Ray on k8s的EP，代表一个Ray集群  - inference：推理的EP，代表一个推理函数实例  - job：Job EP，代表一个可接收Job请求资源组
type EndpointType struct {
	value string
}

type EndpointTypeEnum struct {
	SERVICE   EndpointType
	RAY       EndpointType
	INFERENCE EndpointType
	JOB       EndpointType
}

func GetEndpointTypeEnum() EndpointTypeEnum {
	return EndpointTypeEnum{
		SERVICE: EndpointType{
			value: "service",
		},
		RAY: EndpointType{
			value: "ray",
		},
		INFERENCE: EndpointType{
			value: "inference",
		},
		JOB: EndpointType{
			value: "job",
		},
	}
}

func (c EndpointType) Value() string {
	return c.value
}

func (c EndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointType) UnmarshalJSON(b []byte) error {
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
