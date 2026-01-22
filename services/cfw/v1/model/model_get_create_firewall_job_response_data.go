package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetCreateFirewallJobResponseData struct {

	// **参数解释**： 创建按需防火墙任务ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 任务执行状态，用于展示创建防火墙是否成功。 **取值范围**： - Running：表示任务正在执行。 - Success：表示任务执行成功。 - Failed：表示任务执行失败。
	Status *GetCreateFirewallJobResponseDataStatus `json:"status,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 格式为“yyyy-mm-ddThh:mm:ssZ”，其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 例如： - 2023-10-26T10:30:00Z 表示： 2023年10月26日10点30分00秒，UTC时间 - 2023-10-26T10:30:00+0800 表示： 2023年10月26日10点30分00秒，比UTC时间快8小时，也就是北京时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 格式为“yyyy-mm-ddThh:mm:ssZ”，其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 例如： - 2023-10-26T10:30:00Z 表示： 2023年10月26日10点30分00秒，UTC时间 - 2023-10-26T10:30:00+0800 表示： 2023年10月26日10点30分00秒，比UTC时间快8小时，也就是北京时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o GetCreateFirewallJobResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCreateFirewallJobResponseData struct{}"
	}

	return strings.Join([]string{"GetCreateFirewallJobResponseData", string(data)}, " ")
}

type GetCreateFirewallJobResponseDataStatus struct {
	value string
}

type GetCreateFirewallJobResponseDataStatusEnum struct {
	RUNNING GetCreateFirewallJobResponseDataStatus
	SUCCESS GetCreateFirewallJobResponseDataStatus
	FAILED  GetCreateFirewallJobResponseDataStatus
}

func GetGetCreateFirewallJobResponseDataStatusEnum() GetCreateFirewallJobResponseDataStatusEnum {
	return GetCreateFirewallJobResponseDataStatusEnum{
		RUNNING: GetCreateFirewallJobResponseDataStatus{
			value: "Running",
		},
		SUCCESS: GetCreateFirewallJobResponseDataStatus{
			value: "Success",
		},
		FAILED: GetCreateFirewallJobResponseDataStatus{
			value: "Failed",
		},
	}
}

func (c GetCreateFirewallJobResponseDataStatus) Value() string {
	return c.value
}

func (c GetCreateFirewallJobResponseDataStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetCreateFirewallJobResponseDataStatus) UnmarshalJSON(b []byte) error {
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
