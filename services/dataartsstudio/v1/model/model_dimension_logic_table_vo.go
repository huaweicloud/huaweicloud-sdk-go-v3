package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type DimensionLogicTableVo struct {

	// 表ID
	Id *string `json:"id,omitempty"`

	// 表名称。
	TbName *string `json:"tb_name,omitempty"`

	// 逻辑实体名。
	TbLogicName *string `json:"tb_logic_name,omitempty"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID，填写String类型替代Long类型。
	L3Id *string `json:"l3_id,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 所属维度ID，填写String类型替代Long类型。
	DimensionId *string `json:"dimension_id,omitempty"`

	// 资产责任人。
	Owner *string `json:"owner,omitempty"`

	// 维度类型。 枚举值：   - COMMON: 普通维度   - LOOKUP: 码表维度   - HIERARCHIES: 层级维度
	DimensionType *DimensionLogicTableVoDimensionType `json:"dimension_type,omitempty"`

	// 引用码表ID，填写String类型替代Long类型。
	CodeTableId *string `json:"code_table_id,omitempty"`

	CodeTable *CodeTableVo `json:"code_table,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须。
	QueueName *string `json:"queue_name,omitempty"`

	// 数据连接ID。
	DwId *string `json:"dw_id,omitempty"`

	// 是否是逆向的。
	Reversed *bool `json:"reversed,omitempty"`

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

	// 库名。
	DbName *string `json:"db_name,omitempty"`

	// 数据表ID，只读。
	TbId *string `json:"tb_id,omitempty"`

	// DWS类型需要。
	Schema *string `json:"schema,omitempty"`

	// 表类型。
	TableType *string `json:"table_type,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 表发布后，创建的数据目录技术资产guid，只读，创建和更新时无需填写。
	TbGuid *string `json:"tb_guid,omitempty"`

	// 表发布后，创建的数据目录业务资产guid，只读，创建和更新时无需填写。
	TbLogicGuid *string `json:"tb_logic_guid,omitempty"`

	// 关联维度名称，只读。
	DimensionName *string `json:"dimension_name,omitempty"`

	// 字段属性。
	Attributes *[]DimensionLogicTableAttributeVo `json:"attributes,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType *string `json:"dw_type,omitempty"`

	// 数据连接名称，只读，创建和更新时无需填写。
	DwName *string `json:"dw_name,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	DevPhysicalTable *SyncStatusEnum `json:"dev_physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	Materialization *SyncStatusEnum `json:"materialization,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// DISTRIBUTE BY [HASH(column)|REPLICATION]。 枚举值：   - HASH: 对指定的列进行Hash，通过映射，把数据分布到指定DN   - REPLICATION: 表的每一行存在所有数据节点（DN）中，即每个数据节点都有完整的表数据
	Distribute *DimensionLogicTableVoDistribute `json:"distribute,omitempty"`

	// DISTRIBUTE BY HASH column.
	DistributeColumn *string `json:"distribute_column,omitempty"`

	// 质量ID，填写String类型替代Long类型。
	QualityId *string `json:"quality_id,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`

	// 外表路径
	ObsLocation *string `json:"obs_location,omitempty"`

	// 其他配置
	Configs *string `json:"configs,omitempty"`

	// 开发环境版本，填写String类型替代Long类型。
	DevVersion *string `json:"dev_version,omitempty"`

	// 生产环境版本，填写String类型替代Long类型
	ProdVersion *string `json:"prod_version,omitempty"`

	// 开发环境版本名称
	DevVersionName *string `json:"dev_version_name,omitempty"`

	// 生产环境版本名称
	ProdVersionName *string `json:"prod_version_name,omitempty"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`
}

func (o DimensionLogicTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionLogicTableVo struct{}"
	}

	return strings.Join([]string{"DimensionLogicTableVo", string(data)}, " ")
}

type DimensionLogicTableVoDimensionType struct {
	value string
}

type DimensionLogicTableVoDimensionTypeEnum struct {
	COMMON      DimensionLogicTableVoDimensionType
	LOOKUP      DimensionLogicTableVoDimensionType
	HIERARCHIES DimensionLogicTableVoDimensionType
}

func GetDimensionLogicTableVoDimensionTypeEnum() DimensionLogicTableVoDimensionTypeEnum {
	return DimensionLogicTableVoDimensionTypeEnum{
		COMMON: DimensionLogicTableVoDimensionType{
			value: "COMMON",
		},
		LOOKUP: DimensionLogicTableVoDimensionType{
			value: "LOOKUP",
		},
		HIERARCHIES: DimensionLogicTableVoDimensionType{
			value: "HIERARCHIES",
		},
	}
}

func (c DimensionLogicTableVoDimensionType) Value() string {
	return c.value
}

func (c DimensionLogicTableVoDimensionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionLogicTableVoDimensionType) UnmarshalJSON(b []byte) error {
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

type DimensionLogicTableVoDistribute struct {
	value string
}

type DimensionLogicTableVoDistributeEnum struct {
	HASH        DimensionLogicTableVoDistribute
	REPLICATION DimensionLogicTableVoDistribute
}

func GetDimensionLogicTableVoDistributeEnum() DimensionLogicTableVoDistributeEnum {
	return DimensionLogicTableVoDistributeEnum{
		HASH: DimensionLogicTableVoDistribute{
			value: "HASH",
		},
		REPLICATION: DimensionLogicTableVoDistribute{
			value: "REPLICATION",
		},
	}
}

func (c DimensionLogicTableVoDistribute) Value() string {
	return c.value
}

func (c DimensionLogicTableVoDistribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionLogicTableVoDistribute) UnmarshalJSON(b []byte) error {
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
