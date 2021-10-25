package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowJobStatusResponse struct {
	// Job的状态。 SUCCESS：成功。 RUNNING：运行中。 FAIL：失败。 INIT：正在初始化。

	Status *ShowJobStatusResponseStatus `json:"status,omitempty"`

	Entities *JobEntities `json:"entities,omitempty"`
	// Job ID。

	JobId *string `json:"job_id,omitempty"`
	// Job的类型。createProtectionGroupNoCG：创建保护组。deleteProtectionGroupNoCG：删除保护组。startProtectionGroupNoCG ：保护组开始保护。reprotectProtectionGroupNoCG ：保护组重保护。stopProtectionGroupNoCG ：保护组停止保护。failoverProtectionGroupNoCG  ：保护组故障切换。reverseProtectionGroupNoCG：保护组切换。createProtectedInstanceNoCG：创建保护实例。deleteProtectedInstanceNoCG：删除保护实例。attachReplicationPairNew：保护实例挂载复制对。detachReplicationPairNew：保护实例卸载复制对。addNicNew：保护实例添加网卡。deleteNicNew：保护实例删除网卡。resizeProtectedInstanceNew：保护实例变更规格。createReplicationPairNoCG：创建复制对。deleteReplicationPairNoCG：删除复制对。expandReplicationPairNew：复制对扩容。createDisasterRecoveryDrill：创建容灾演练。deleteDisasterRecoveryDrill：删除容灾演练。

	JobType *string `json:"job_type,omitempty"`
	// 开始时间。默认格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSSZ\"，例：\"2019-04-01T12:00:00.000Z\"。

	BeginTime *string `json:"begin_time,omitempty"`
	// 结束时间。默认格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSSZ\"，例：\"2019-04-01T12:00:00.000Z\"。

	EndTime *string `json:"end_time,omitempty"`
	// Job执行失败时的错误码。

	ErrorCode *string `json:"error_code,omitempty"`
	// Job执行失败时的错误原因。

	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}

type ShowJobStatusResponseStatus struct {
	value string
}

type ShowJobStatusResponseStatusEnum struct {
	SUCCESS ShowJobStatusResponseStatus
	RUNNING ShowJobStatusResponseStatus
	FAIL    ShowJobStatusResponseStatus
	INIT    ShowJobStatusResponseStatus
}

func GetShowJobStatusResponseStatusEnum() ShowJobStatusResponseStatusEnum {
	return ShowJobStatusResponseStatusEnum{
		SUCCESS: ShowJobStatusResponseStatus{
			value: "SUCCESS",
		},
		RUNNING: ShowJobStatusResponseStatus{
			value: "RUNNING",
		},
		FAIL: ShowJobStatusResponseStatus{
			value: "FAIL",
		},
		INIT: ShowJobStatusResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
