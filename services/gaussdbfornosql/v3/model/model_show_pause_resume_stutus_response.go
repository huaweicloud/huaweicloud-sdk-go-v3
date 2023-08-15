package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPauseResumeStutusResponse Response Object
type ShowPauseResumeStutusResponse struct {

	// 主实例id
	MasterInstanceId *string `json:"master_instance_id,omitempty"`

	// 备实例id
	SlaveInstanceId *string `json:"slave_instance_id,omitempty"`

	// 容灾实例数据同步状态 - NA：实例尚未搭建容灾关系 - NEW：尚未启动的数据同步状态 - SYNCING：数据同步正常进行中 - SUSPENDING：正在暂停数据同步 - SUSPENDED：数据同步已暂停 - RECOVERYING：正在恢复数据同步
	Status *ShowPauseResumeStutusResponseStatus `json:"status,omitempty"`

	DataSyncIndicators *NoSqlDrDateSyncIndicators `json:"data_sync_indicators,omitempty"`

	// 切换或倒换RPO和RTO值，仅当请求实例id为主实例时有值
	RtoAndRpoIndicators *[]NoSqlDrRpoAndRto `json:"rto_and_rpo_indicators,omitempty"`
	HttpStatusCode      int                 `json:"-"`
}

func (o ShowPauseResumeStutusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPauseResumeStutusResponse struct{}"
	}

	return strings.Join([]string{"ShowPauseResumeStutusResponse", string(data)}, " ")
}

type ShowPauseResumeStutusResponseStatus struct {
	value string
}

type ShowPauseResumeStutusResponseStatusEnum struct {
	NA          ShowPauseResumeStutusResponseStatus
	NEW         ShowPauseResumeStutusResponseStatus
	SYNCING     ShowPauseResumeStutusResponseStatus
	SUSPENDING  ShowPauseResumeStutusResponseStatus
	SUSPENDED   ShowPauseResumeStutusResponseStatus
	RECOVERYING ShowPauseResumeStutusResponseStatus
}

func GetShowPauseResumeStutusResponseStatusEnum() ShowPauseResumeStutusResponseStatusEnum {
	return ShowPauseResumeStutusResponseStatusEnum{
		NA: ShowPauseResumeStutusResponseStatus{
			value: "NA",
		},
		NEW: ShowPauseResumeStutusResponseStatus{
			value: "NEW",
		},
		SYNCING: ShowPauseResumeStutusResponseStatus{
			value: "SYNCING",
		},
		SUSPENDING: ShowPauseResumeStutusResponseStatus{
			value: "SUSPENDING",
		},
		SUSPENDED: ShowPauseResumeStutusResponseStatus{
			value: "SUSPENDED",
		},
		RECOVERYING: ShowPauseResumeStutusResponseStatus{
			value: "RECOVERYING",
		},
	}
}

func (c ShowPauseResumeStutusResponseStatus) Value() string {
	return c.value
}

func (c ShowPauseResumeStutusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPauseResumeStutusResponseStatus) UnmarshalJSON(b []byte) error {
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
