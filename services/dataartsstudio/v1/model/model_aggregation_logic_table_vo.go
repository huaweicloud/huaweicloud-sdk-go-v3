package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AggregationLogicTableVo 汇总表值对象（VO）。
type AggregationLogicTableVo struct {

	// 汇总表的唯一系统ID，更新时必填，创建时不须填写，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 汇总表英文名称，对应实际的物理表名。
	TbName string `json:"tb_name"`

	// 汇总表的中文名，用于展示使用。
	TbLogicName string `json:"tb_logic_name"`

	// 主题域分组ID，只读，创建和更新时无需填写，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 汇总表所属主题的ID，必填，填写String类型替代Long类型。
	L3Id string `json:"l3_id"`

	// 汇总表描述信息。
	Description *string `json:"description,omitempty"`

	// 汇总表的资产责任人。
	Owner string `json:"owner"`

	SecretType *SecretTypeEnum `json:"secret_type,omitempty"`

	ApplyBg *ApplyBgEnum `json:"apply_bg,omitempty"`

	// 汇总表的创建人，只读，创建和更新时无需填写。
	CreateBy *string `json:"create_by,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须填写。
	QueueName *string `json:"queue_name,omitempty"`

	// 汇总表所在的数据连接ID，为32位十六进制数字。
	DwId string `json:"dw_id"`

	// 汇总表所在的数据库名。
	DbName string `json:"db_name"`

	// 汇总表创建的表ID，是服务内部ID，只读，创建和更新时无需填写
	TbId *string `json:"tb_id,omitempty"`

	// 汇总表所在的Schema名，DWS类型必须填写。
	Schema *string `json:"schema,omitempty"`

	// 数据连接名称，只读，创建和更新时无需填写。
	DwName *string `json:"dw_name,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 表发布后，创建的数据目录技术资产guid，只读，创建和更新时无需填写。
	TbGuid *string `json:"tb_guid,omitempty"`

	// 表发布后，创建的数据目录业务资产guid，只读，创建和更新时无需填写。
	TbLogicGuid *string `json:"tb_logic_guid,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType string `json:"dw_type"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 创建时间，只读，创建和更新时无需填写。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，创建和更新时无需填写。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 颗粒度ID。
	DimensionGroup *string `json:"dimension_group,omitempty"`

	// 颗粒度名称，只读。
	GroupName *string `json:"group_name,omitempty"`

	// 颗粒度编码，只读。
	GroupCode *string `json:"group_code,omitempty"`

	TimePeriod *AggregationLogicTableAttributeVo `json:"time_period,omitempty"`

	// 汇总表属性信息，依据attribute_type判断类型。
	TableAttributes *[]AggregationLogicTableAttributeVo `json:"table_attributes,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	DevPhysicalTable *SyncStatusEnum `json:"dev_physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	DlfTask *SyncStatusEnum `json:"dlf_task,omitempty"`

	PublishToDlm *SyncStatusEnum `json:"publish_to_dlm,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// DISTRIBUTE BY [HASH(column)|REPLICATION]。 枚举值：   - HASH: 对指定的列进行Hash，通过映射，把数据分布到指定DN   - REPLICATION: 表的每一行存在所有数据节点（DN）中，即每个数据节点都有完整的表数据
	Distribute *AggregationLogicTableVoDistribute `json:"distribute,omitempty"`

	// DISTRIBUTE BY HASH column.
	DistributeColumn *string `json:"distribute_column,omitempty"`

	// DWS数据压缩等级，列压缩等级为no/low/middle/high，行压缩等级为no/yes。 枚举值：   - \"NO\": 不压缩   - \"YES\": 压缩   - \"LOW\": 低等级压缩   - \"MIDDLE\": 中等级压缩   - \"HIGH\": 高等级压缩
	Compression *AggregationLogicTableVoCompression `json:"compression,omitempty"`

	// 外表路径。
	ObsLocation *string `json:"obs_location,omitempty"`

	// 版本字段。
	PreCombineField *string `json:"pre_combine_field,omitempty"`

	// 表类型。
	TableType *string `json:"table_type,omitempty"`

	// DLF作业ID。
	DlfTaskId *string `json:"dlf_task_id,omitempty"`

	// 质量ID，填写String类型替代Long类型。
	QualityId *string `json:"quality_id,omitempty"`

	// 是否是逆向的，只读。
	Reversed *bool `json:"reversed,omitempty"`

	// 为2时，表示汇总表是汇总生成的，只读。
	TableVersion *int32 `json:"table_version,omitempty"`

	// 分区表达式。
	PartitionConf *string `json:"partition_conf,omitempty"`

	// 异常数据输出开关。
	DirtyOutSwitch *bool `json:"dirty_out_switch,omitempty"`

	// 异常数据输出库。
	DirtyOutDatabase *string `json:"dirty_out_database,omitempty"`

	// 异常表前缀。
	DirtyOutPrefix *string `json:"dirty_out_prefix,omitempty"`

	// 异常表后缀。
	DirtyOutSuffix *string `json:"dirty_out_suffix,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 其他配置。
	Configs *string `json:"configs,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`

	// API ID。
	ApiId *string `json:"api_id,omitempty"`

	// 汇总表绑定的SQL。
	Sql *string `json:"sql,omitempty"`

	// 开发环境版本，填写String类型替代Long类型。
	DevVersion *string `json:"dev_version,omitempty"`

	// 生产环境版本，填写String类型替代Long类型。
	ProdVersion *string `json:"prod_version,omitempty"`

	// 开发环境版本名称
	DevVersionName *string `json:"dev_version_name,omitempty"`

	// 生产环境版本名称
	ProdVersionName *string `json:"prod_version_name,omitempty"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`
}

func (o AggregationLogicTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationLogicTableVo struct{}"
	}

	return strings.Join([]string{"AggregationLogicTableVo", string(data)}, " ")
}

type AggregationLogicTableVoDistribute struct {
	value string
}

type AggregationLogicTableVoDistributeEnum struct {
	HASH        AggregationLogicTableVoDistribute
	REPLICATION AggregationLogicTableVoDistribute
}

func GetAggregationLogicTableVoDistributeEnum() AggregationLogicTableVoDistributeEnum {
	return AggregationLogicTableVoDistributeEnum{
		HASH: AggregationLogicTableVoDistribute{
			value: "HASH",
		},
		REPLICATION: AggregationLogicTableVoDistribute{
			value: "REPLICATION",
		},
	}
}

func (c AggregationLogicTableVoDistribute) Value() string {
	return c.value
}

func (c AggregationLogicTableVoDistribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AggregationLogicTableVoDistribute) UnmarshalJSON(b []byte) error {
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

type AggregationLogicTableVoCompression struct {
	value string
}

type AggregationLogicTableVoCompressionEnum struct {
	NO     AggregationLogicTableVoCompression
	YES    AggregationLogicTableVoCompression
	LOW    AggregationLogicTableVoCompression
	MIDDLE AggregationLogicTableVoCompression
	HIGH   AggregationLogicTableVoCompression
}

func GetAggregationLogicTableVoCompressionEnum() AggregationLogicTableVoCompressionEnum {
	return AggregationLogicTableVoCompressionEnum{
		NO: AggregationLogicTableVoCompression{
			value: "NO",
		},
		YES: AggregationLogicTableVoCompression{
			value: "YES",
		},
		LOW: AggregationLogicTableVoCompression{
			value: "LOW",
		},
		MIDDLE: AggregationLogicTableVoCompression{
			value: "MIDDLE",
		},
		HIGH: AggregationLogicTableVoCompression{
			value: "HIGH",
		},
	}
}

func (c AggregationLogicTableVoCompression) Value() string {
	return c.value
}

func (c AggregationLogicTableVoCompression) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AggregationLogicTableVoCompression) UnmarshalJSON(b []byte) error {
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
