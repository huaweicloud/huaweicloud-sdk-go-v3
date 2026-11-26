package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AssetOperationRecordRsp **参数解释**：资产操作记录。
type AssetOperationRecordRsp struct {

	// **参数解释**：操作记录ID。 **取值范围**：不涉及。
	RecordId *string `json:"record_id,omitempty"`

	// **参数解释**：华为云项目ID，获取方法请参见[获取项目ID和名称](modelarts_03_0147.xml)。 **取值范围**：不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：工作空间ID，[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。]未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：作业类型。 **取值范围**：枚举值: - PRETRAIN: 预训练 - SFT: 全量微调 - RFT: 强化微调 - RLHF: 强化学习 - QUANTIZATION: 量化 - MANUAL-EVALUATION: 人工评测 - AUTO-EVALUATION: 自动评测 - ONLINE-DEPLOY: 在线部署 - EDGE-DEPLOY: 边缘部署 - LORA: LoRA微调 - QLORA: Q-LoRA微调 - DATA-CLEANING: 数据清洗 - DATA-IMPORT: 数据导入 - DATA-ANNOTATION: 数据标注
	JobType AssetOperationRecordRspJobType `json:"job_type"`

	// **参数解释**：关联任务ID。 **取值范围**：不涉及。
	JobId string `json:"job_id"`

	// **参数解释**：关联任务名称。 **取值范围**：不涉及。
	JobName string `json:"job_name"`

	// **参数解释**：关联任务状态。 **取值范围**：枚举值: - INIT: 初始化 - RUNNING: 运行中 - STOPPED: 已停止 - FAILED: 失败 - SUCCEED: 成功 - DELETED: 已删除
	JobStatus AssetOperationRecordRspJobStatus `json:"job_status"`

	// **参数解释**：操作用户ID。 **取值范围**：不涉及。
	UserId string `json:"user_id"`

	// **参数解释**：操作用户的用户名。 **取值范围**：不涉及。
	UserName string `json:"user_name"`

	// **参数解释**：资产ID。 **取值范围**：不涉及。
	AssetId string `json:"asset_id"`

	// **参数解释**：来源服务。 **取值范围**：不涉及。
	Service string `json:"service"`

	// **参数解释**：资产编码。 **取值范围**：不涉及。
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**：资产名称，支持英文字母，数字，中划线，下划线，点，空格，中文汉字。 **取值范围**：不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**：资产版本。 **取值范围**：不涉及。
	AssetVersion *string `json:"asset_version,omitempty"`

	// **参数解释**：操作记录创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// **参数解释**：操作记录更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// **参数解释**：任务是否删除（true: 已删除，false: 未删除）。 **取值范围**：不涉及。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// **参数解释**：服务操作详情。 **取值范围**：不涉及。
	Details *string `json:"details,omitempty"`

	// **参数解释**：预留扩展信息字段。 **取值范围**：不涉及。
	RecordExtendMessage *string `json:"record_extend_message,omitempty"`
}

func (o AssetOperationRecordRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetOperationRecordRsp struct{}"
	}

	return strings.Join([]string{"AssetOperationRecordRsp", string(data)}, " ")
}

type AssetOperationRecordRspJobType struct {
	value string
}

type AssetOperationRecordRspJobTypeEnum struct {
	PRETRAIN          AssetOperationRecordRspJobType
	SFT               AssetOperationRecordRspJobType
	RFT               AssetOperationRecordRspJobType
	RLHF              AssetOperationRecordRspJobType
	QUANTIZATION      AssetOperationRecordRspJobType
	MANUAL_EVALUATION AssetOperationRecordRspJobType
	AUTO_EVALUATION   AssetOperationRecordRspJobType
	ONLINE_DEPLOY     AssetOperationRecordRspJobType
	EDGE_DEPLOY       AssetOperationRecordRspJobType
	LORA              AssetOperationRecordRspJobType
	QLORA             AssetOperationRecordRspJobType
	DATA_CLEANING     AssetOperationRecordRspJobType
	DATA_IMPORT       AssetOperationRecordRspJobType
	DATA_ANNOTATION   AssetOperationRecordRspJobType
}

func GetAssetOperationRecordRspJobTypeEnum() AssetOperationRecordRspJobTypeEnum {
	return AssetOperationRecordRspJobTypeEnum{
		PRETRAIN: AssetOperationRecordRspJobType{
			value: "PRETRAIN",
		},
		SFT: AssetOperationRecordRspJobType{
			value: "SFT",
		},
		RFT: AssetOperationRecordRspJobType{
			value: "RFT",
		},
		RLHF: AssetOperationRecordRspJobType{
			value: "RLHF",
		},
		QUANTIZATION: AssetOperationRecordRspJobType{
			value: "QUANTIZATION",
		},
		MANUAL_EVALUATION: AssetOperationRecordRspJobType{
			value: "MANUAL-EVALUATION",
		},
		AUTO_EVALUATION: AssetOperationRecordRspJobType{
			value: "AUTO-EVALUATION",
		},
		ONLINE_DEPLOY: AssetOperationRecordRspJobType{
			value: "ONLINE-DEPLOY",
		},
		EDGE_DEPLOY: AssetOperationRecordRspJobType{
			value: "EDGE-DEPLOY",
		},
		LORA: AssetOperationRecordRspJobType{
			value: "LORA",
		},
		QLORA: AssetOperationRecordRspJobType{
			value: "QLORA",
		},
		DATA_CLEANING: AssetOperationRecordRspJobType{
			value: "DATA-CLEANING",
		},
		DATA_IMPORT: AssetOperationRecordRspJobType{
			value: "DATA-IMPORT",
		},
		DATA_ANNOTATION: AssetOperationRecordRspJobType{
			value: "DATA-ANNOTATION",
		},
	}
}

func (c AssetOperationRecordRspJobType) Value() string {
	return c.value
}

func (c AssetOperationRecordRspJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetOperationRecordRspJobType) UnmarshalJSON(b []byte) error {
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

type AssetOperationRecordRspJobStatus struct {
	value string
}

type AssetOperationRecordRspJobStatusEnum struct {
	INIT    AssetOperationRecordRspJobStatus
	RUNNING AssetOperationRecordRspJobStatus
	STOPPED AssetOperationRecordRspJobStatus
	FAILED  AssetOperationRecordRspJobStatus
	SUCCEED AssetOperationRecordRspJobStatus
	DELETED AssetOperationRecordRspJobStatus
}

func GetAssetOperationRecordRspJobStatusEnum() AssetOperationRecordRspJobStatusEnum {
	return AssetOperationRecordRspJobStatusEnum{
		INIT: AssetOperationRecordRspJobStatus{
			value: "INIT",
		},
		RUNNING: AssetOperationRecordRspJobStatus{
			value: "RUNNING",
		},
		STOPPED: AssetOperationRecordRspJobStatus{
			value: "STOPPED",
		},
		FAILED: AssetOperationRecordRspJobStatus{
			value: "FAILED",
		},
		SUCCEED: AssetOperationRecordRspJobStatus{
			value: "SUCCEED",
		},
		DELETED: AssetOperationRecordRspJobStatus{
			value: "DELETED",
		},
	}
}

func (c AssetOperationRecordRspJobStatus) Value() string {
	return c.value
}

func (c AssetOperationRecordRspJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetOperationRecordRspJobStatus) UnmarshalJSON(b []byte) error {
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
