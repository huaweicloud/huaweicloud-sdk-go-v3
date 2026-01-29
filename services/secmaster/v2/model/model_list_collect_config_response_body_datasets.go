package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListCollectConfigResponseBodyDatasets struct {

	// 接入账号总数量
	AccountAllNum float32 `json:"account_all_num,omitempty"`

	// 成功接入账号数量
	AccountSuccessfulNum float32 `json:"account_successful_num,omitempty"`

	// 账号列表
	Accounts *[]ListCollectConfigResponseBodyAccounts `json:"accounts,omitempty"`

	// 自动转告警开关
	Alert *bool `json:"alert,omitempty"`

	// 是否接入已纳管的全量账号
	AllAccounts *bool `json:"all_accounts,omitempty"`

	// 能否开自动转告警
	AllowAlert *bool `json:"allow_alert,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 开启状态
	Enable *string `json:"enable,omitempty"`

	// 上次活跃时间
	LastActiveTime float32 `json:"last_active_time,omitempty"`

	// 限制
	Limit *string `json:"limit,omitempty"`

	// 新账号自动接入开关
	NewAccountAutoAccess *bool `json:"new_account_auto_access,omitempty"`

	// 日志的接入状态,可能的值为FAIL,DEFAULT,CREATING,SUCCESS,FAIL表示失败,DEFAULT表示待接入,CREATING表示接入中,SUCCESS表示成功
	ProcessStatus *ListCollectConfigResponseBodyDatasetsProcessStatus `json:"process_status,omitempty"`

	Reference *ListCollectConfigResponseBodyReference `json:"reference,omitempty"`

	// 所属区域
	RegionId *string `json:"region_id,omitempty"`

	// 错误信息
	SinkMsg *string `json:"sink_msg,omitempty"`

	// 日志ID
	SourceId float32 `json:"source_id,omitempty"`

	// 日志名称
	SourceName *string `json:"source_name,omitempty"`

	Target *ListCollectConfigResponseBodyTarget `json:"target,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ListCollectConfigResponseBodyDatasets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyDatasets struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyDatasets", string(data)}, " ")
}

type ListCollectConfigResponseBodyDatasetsProcessStatus struct {
	value string
}

type ListCollectConfigResponseBodyDatasetsProcessStatusEnum struct {
	FAIL     ListCollectConfigResponseBodyDatasetsProcessStatus
	DEFAULT  ListCollectConfigResponseBodyDatasetsProcessStatus
	CREATING ListCollectConfigResponseBodyDatasetsProcessStatus
	SUCCESS  ListCollectConfigResponseBodyDatasetsProcessStatus
}

func GetListCollectConfigResponseBodyDatasetsProcessStatusEnum() ListCollectConfigResponseBodyDatasetsProcessStatusEnum {
	return ListCollectConfigResponseBodyDatasetsProcessStatusEnum{
		FAIL: ListCollectConfigResponseBodyDatasetsProcessStatus{
			value: "FAIL",
		},
		DEFAULT: ListCollectConfigResponseBodyDatasetsProcessStatus{
			value: "DEFAULT",
		},
		CREATING: ListCollectConfigResponseBodyDatasetsProcessStatus{
			value: "CREATING",
		},
		SUCCESS: ListCollectConfigResponseBodyDatasetsProcessStatus{
			value: "SUCCESS",
		},
	}
}

func (c ListCollectConfigResponseBodyDatasetsProcessStatus) Value() string {
	return c.value
}

func (c ListCollectConfigResponseBodyDatasetsProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectConfigResponseBodyDatasetsProcessStatus) UnmarshalJSON(b []byte) error {
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
