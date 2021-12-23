package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResizeInstanceReq struct {
	// 规格变更后的规格ID。 请参考[查询可扩容规格列表](https://support.huaweicloud.com/api-rabbitmq/ResizeInstance.html)接口返回的数据。

	NewSpecCode ResizeInstanceReqNewSpecCode `json:"new_spec_code"`
	// 规格变更后的消息存储空间，单位：GB。 请参考[查询可扩容规格列表](https://support.huaweicloud.com/api-rabbitmq/ResizeInstance.html)接口返回的数据。

	NewStorageSpace int32 `json:"new_storage_space"`
}

func (o ResizeInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceReq struct{}"
	}

	return strings.Join([]string{"ResizeInstanceReq", string(data)}, " ")
}

type ResizeInstanceReqNewSpecCode struct {
	value string
}

type ResizeInstanceReqNewSpecCodeEnum struct {
	DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_3 ResizeInstanceReqNewSpecCode
	DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_5 ResizeInstanceReqNewSpecCode
	DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_7 ResizeInstanceReqNewSpecCode
}

func GetResizeInstanceReqNewSpecCodeEnum() ResizeInstanceReqNewSpecCodeEnum {
	return ResizeInstanceReqNewSpecCodeEnum{
		DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_3: ResizeInstanceReqNewSpecCode{
			value: "dms.instance.rabbitmq.cluster.c3.4u8g.3",
		},
		DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_5: ResizeInstanceReqNewSpecCode{
			value: "dms.instance.rabbitmq.cluster.c3.4u8g.5",
		},
		DMS_INSTANCE_RABBITMQ_CLUSTER_C3_4U8G_7: ResizeInstanceReqNewSpecCode{
			value: "dms.instance.rabbitmq.cluster.c3.4u8g.7",
		},
	}
}

func (c ResizeInstanceReqNewSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceReqNewSpecCode) UnmarshalJSON(b []byte) error {
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
