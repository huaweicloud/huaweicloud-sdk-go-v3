package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOfflineKeyAnalysisTaskResponse Response Object
type ShowOfflineKeyAnalysisTaskResponse struct {

	// **参数解释**： 任务执行记录ID（此ID对应“查询离线全量key分析详情”参数中的任务ID）。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**： 备份ID。 **取值范围**： 不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释**： 分片名称。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 节点ipv4地址。 **取值范围**： 不涉及。
	NodeIp *string `json:"node_ip,omitempty"`

	// **参数解释**： 节点ipv6地址。 **取值范围**： 不涉及。
	NodeIpv6 *string `json:"node_ipv6,omitempty"`

	// **参数解释**： 节点类型。 **取值范围**： master：主节点 slave：从节点
	NodeType *string `json:"node_type,omitempty"`

	// **参数解释**： 分析类型。 **取值范围**： new_backup：新建备份并分析。 old_backup：历史备份文件分析。
	AnalysisType *ShowOfflineKeyAnalysisTaskResponseAnalysisType `json:"analysis_type,omitempty"`

	// **参数解释**： 分析任务开始时间。 **取值范围**： 不涉及。
	StartedAt *string `json:"started_at,omitempty"`

	// **参数解释**： 分析任务结束时间。 **取值范围**： 不涉及。
	FinishedAt *string `json:"finished_at,omitempty"`

	// **参数解释**： 前缀Key列表。
	LargestKeyPrefixes *[]LargestKeyPrefix `json:"largest_key_prefixes,omitempty"`

	// **参数解释**： 大key列表。
	LargestKeys *[]LargestKey `json:"largest_keys,omitempty"`

	// **参数解释**： Key的总大小，单位：Bytes。 **取值范围**： 不涉及。
	TotalBytes *int32 `json:"total_bytes,omitempty"`

	// **参数解释**： Key的总数。 **取值范围**： 不涉及。
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 每种类型key的总大小，单位：Bytes。 **取值范围**： 不涉及。
	TypeBytes *[]KeyTypeByte `json:"type_bytes,omitempty"`

	// **参数解释**： 每种类型key总数。 **取值范围**： 不涉及。
	TypeNum        *[]KeyTypeNum `json:"type_num,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowOfflineKeyAnalysisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOfflineKeyAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowOfflineKeyAnalysisTaskResponse", string(data)}, " ")
}

type ShowOfflineKeyAnalysisTaskResponseAnalysisType struct {
	value string
}

type ShowOfflineKeyAnalysisTaskResponseAnalysisTypeEnum struct {
	NEW_BACKUP ShowOfflineKeyAnalysisTaskResponseAnalysisType
	OLD_BACKUP ShowOfflineKeyAnalysisTaskResponseAnalysisType
}

func GetShowOfflineKeyAnalysisTaskResponseAnalysisTypeEnum() ShowOfflineKeyAnalysisTaskResponseAnalysisTypeEnum {
	return ShowOfflineKeyAnalysisTaskResponseAnalysisTypeEnum{
		NEW_BACKUP: ShowOfflineKeyAnalysisTaskResponseAnalysisType{
			value: "new_backup",
		},
		OLD_BACKUP: ShowOfflineKeyAnalysisTaskResponseAnalysisType{
			value: "old_backup",
		},
	}
}

func (c ShowOfflineKeyAnalysisTaskResponseAnalysisType) Value() string {
	return c.value
}

func (c ShowOfflineKeyAnalysisTaskResponseAnalysisType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOfflineKeyAnalysisTaskResponseAnalysisType) UnmarshalJSON(b []byte) error {
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
