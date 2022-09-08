package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceStatusResponse struct {

	// 服务链接
	ServerUrl *string `json:"server_url,omitempty"`

	// 实例状态。 - DELETED 已删除 - DELETE_FAILED 删除失败 - DELETING 删除中 - READY 热实例就绪状态 - RUNNING 正在运行 - STARTING 正在启动 - STOPPED 已停止 - STOPPING 停止中 - UPDATE 更新Schdule信息 - WAITING 热实例创建初始态
	Status *InstanceStatusResponseStatus `json:"status,omitempty"`
}

func (o InstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"InstanceStatusResponse", string(data)}, " ")
}

type InstanceStatusResponseStatus struct {
	value string
}

type InstanceStatusResponseStatusEnum struct {
	DELETED       InstanceStatusResponseStatus
	DELETE_FAILED InstanceStatusResponseStatus
	DELETING      InstanceStatusResponseStatus
	READY         InstanceStatusResponseStatus
	RUNNING       InstanceStatusResponseStatus
	STARTING      InstanceStatusResponseStatus
	STOPPED       InstanceStatusResponseStatus
	STOPPING      InstanceStatusResponseStatus
	UPDATE        InstanceStatusResponseStatus
	WAITING       InstanceStatusResponseStatus
}

func GetInstanceStatusResponseStatusEnum() InstanceStatusResponseStatusEnum {
	return InstanceStatusResponseStatusEnum{
		DELETED: InstanceStatusResponseStatus{
			value: "DELETED",
		},
		DELETE_FAILED: InstanceStatusResponseStatus{
			value: "DELETE_FAILED",
		},
		DELETING: InstanceStatusResponseStatus{
			value: "DELETING",
		},
		READY: InstanceStatusResponseStatus{
			value: "READY",
		},
		RUNNING: InstanceStatusResponseStatus{
			value: "RUNNING",
		},
		STARTING: InstanceStatusResponseStatus{
			value: "STARTING",
		},
		STOPPED: InstanceStatusResponseStatus{
			value: "STOPPED",
		},
		STOPPING: InstanceStatusResponseStatus{
			value: "STOPPING",
		},
		UPDATE: InstanceStatusResponseStatus{
			value: "UPDATE",
		},
		WAITING: InstanceStatusResponseStatus{
			value: "WAITING",
		},
	}
}

func (c InstanceStatusResponseStatus) Value() string {
	return c.value
}

func (c InstanceStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
