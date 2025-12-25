package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChannelOperationStatus **参数解释**: 采集通道部署的进度 - READY 等待部署 - START 开始部署 - SUCCESS 部署成功 - FAIL 部署失败  **约束限制** 不涉及 **取值范围**: - READY - START - SUCCESS - FAIL  **默认值** 不涉及
type ChannelOperationStatus struct {
	value string
}

type ChannelOperationStatusEnum struct {
	READY   ChannelOperationStatus
	START   ChannelOperationStatus
	SUCCESS ChannelOperationStatus
	FAIL    ChannelOperationStatus
}

func GetChannelOperationStatusEnum() ChannelOperationStatusEnum {
	return ChannelOperationStatusEnum{
		READY: ChannelOperationStatus{
			value: "READY",
		},
		START: ChannelOperationStatus{
			value: "START",
		},
		SUCCESS: ChannelOperationStatus{
			value: "SUCCESS",
		},
		FAIL: ChannelOperationStatus{
			value: "FAIL",
		},
	}
}

func (c ChannelOperationStatus) Value() string {
	return c.value
}

func (c ChannelOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChannelOperationStatus) UnmarshalJSON(b []byte) error {
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
