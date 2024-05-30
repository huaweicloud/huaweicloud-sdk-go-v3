package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type TableModelUpdateVo struct {

	// 编码，填写String类型替代Long类型。
	Id string `json:"id"`

	// 所属关系建模的模型ID，填写String类型替代Long类型。
	ModelId string `json:"model_id"`

	// 父表ID，填写String类型替代Long类型。
	ParentTableId *string `json:"parent_table_id,omitempty"`

	// 父表名称，只读。
	ParentTableName *string `json:"parent_table_name,omitempty"`

	// 父表编码，只读。
	ParentTableCode *string `json:"parent_table_code,omitempty"`

	Model *WorkspaceVo `json:"model,omitempty"`

	// 数据格式。
	DataFormat *string `json:"data_format,omitempty"`

	// obs桶。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// obs路径。
	ObsLocation *string `json:"obs_location,omitempty"`

	// 其他配置。
	Configs *string `json:"configs,omitempty"`

	// 表类型，只读。
	TableType *string `json:"table_type,omitempty"`

	// 责任人。
	Owner *string `json:"owner,omitempty"`

	// 表名。
	TbName string `json:"tb_name"`

	// 数据连接ID。
	DwId *string `json:"dw_id,omitempty"`

	// 数据库名。
	DbName *string `json:"db_name,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须。
	QueueName *string `json:"queue_name,omitempty"`

	// DWS类型需要。
	Schema *string `json:"schema,omitempty"`

	// 扩展信息。
	ExtendInfo *string `json:"extend_info,omitempty"`

	// 表物化后的guid，只读。
	TbGuid *string `json:"tb_guid,omitempty"`

	// 数据表ID，只读。
	TbId *string `json:"tb_id,omitempty"`

	// 逻辑实体名。
	LogicTbName string `json:"logic_tb_name"`

	// 逻辑实体的guid，只读。
	LogicTbGuid *string `json:"logic_tb_guid,omitempty"`

	// 描述。
	Description string `json:"description"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 逻辑实体的ID，填写String类型替代Long类型。
	LogicTbId *string `json:"logic_tb_id,omitempty"`

	// 归属的业务分类的ID，填写String类型替代Long类型。
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 归属的业务分类的路径，格式： {\"l1Id\":\"958408897973161984\",\"l2Id\":\"958408897973161985\",\"l3Id\":\"958408897973161986\"}。
	CatalogPath *string `json:"catalog_path,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人，只读。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 表标签，只读。
	Tags *[]TagVo `json:"tags,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 表属性信息。
	Attributes []TableModelAttributeVo `json:"attributes"`

	// 表映射信息。
	Mappings *[]TableMappingVo `json:"mappings,omitempty"`

	// 关系。
	Relations *[]RelationVo `json:"relations,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType string `json:"dw_type"`

	// 数据连接名称，只读，创建和更新时无需填写。
	DwName *string `json:"dw_name,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID，只读，填写String类型替代Long类型。
	L3Id *string `json:"l3_id,omitempty"`

	// 分区表达式
	PartitionConf *string `json:"partition_conf,omitempty"`

	// DLF作业ID。
	DlfTaskId *string `json:"dlf_task_id,omitempty"`

	// 是否使用最新分区。
	UseRecentlyPartition *bool `json:"use_recently_partition,omitempty"`

	// 是否是逆向的。
	Reversed *bool `json:"reversed,omitempty"`

	// 异常数据输出开关。
	DirtyOutSwitch *bool `json:"dirty_out_switch,omitempty"`

	// 异常数据输出库。
	DirtyOutDatabase *string `json:"dirty_out_database,omitempty"`

	// 异常表前缀。
	DirtyOutPrefix *string `json:"dirty_out_prefix,omitempty"`

	// 异常表后缀。
	DirtyOutSuffix *string `json:"dirty_out_suffix,omitempty"`

	// 质量责任人。
	QualityOwner *string `json:"quality_owner,omitempty"`

	// 质量ID，填写String类型替代Long类型。
	QualityId *string `json:"quality_id,omitempty"`

	// DISTRIBUTE BY [HASH(column)|REPLICATION]。 枚举值：   - HASH: 对指定的列进行Hash，通过映射，把数据分布到指定DN   - REPLICATION: 表的每一行存在所有数据节点（DN）中，即每个数据节点都有完整的表数据
	Distribute *TableModelUpdateVoDistribute `json:"distribute,omitempty"`

	// DISTRIBUTE BY HASH column.
	DistributeColumn *string `json:"distribute_column,omitempty"`

	// 是否分区表，只读。
	IsPartition *bool `json:"is_partition,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	DevPhysicalTable *SyncStatusEnum `json:"dev_physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`

	// 开发环境版本，填写String类型替代Long类型。
	DevVersion *string `json:"dev_version,omitempty"`

	// 生产环境版本，填写String类型替代Long类型。
	ProdVersion *string `json:"prod_version,omitempty"`

	// 开发环境版本名称
	DevVersionName *string `json:"dev_version_name,omitempty"`

	// 生产环境版本名称
	ProdVersionName *string `json:"prod_version_name,omitempty"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`

	// 是否关联了物理表
	HasRelatedPhysicalTable *bool `json:"has_related_physical_table,omitempty"`

	// 是否关联了逻辑实体
	HasRelatedLogicTable *bool `json:"has_related_logic_table,omitempty"`
}

func (o TableModelUpdateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableModelUpdateVo struct{}"
	}

	return strings.Join([]string{"TableModelUpdateVo", string(data)}, " ")
}

type TableModelUpdateVoDistribute struct {
	value string
}

type TableModelUpdateVoDistributeEnum struct {
	HASH        TableModelUpdateVoDistribute
	REPLICATION TableModelUpdateVoDistribute
}

func GetTableModelUpdateVoDistributeEnum() TableModelUpdateVoDistributeEnum {
	return TableModelUpdateVoDistributeEnum{
		HASH: TableModelUpdateVoDistribute{
			value: "HASH",
		},
		REPLICATION: TableModelUpdateVoDistribute{
			value: "REPLICATION",
		},
	}
}

func (c TableModelUpdateVoDistribute) Value() string {
	return c.value
}

func (c TableModelUpdateVoDistribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableModelUpdateVoDistribute) UnmarshalJSON(b []byte) error {
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
