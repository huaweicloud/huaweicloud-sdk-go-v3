package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllTableVo 总览表。
type AllTableVo struct {

	// l1的ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// l1名称。
	Name *string `json:"name,omitempty"`

	// 表发布后对应的逻辑实体guid。
	TbLogicGuid *string `json:"tb_logic_guid,omitempty"`

	// 质量ID。
	QualityId *string `json:"quality_id,omitempty"`

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

	// 表发布后对应的物理表guid。
	TbGuid *string `json:"tb_guid,omitempty"`

	// 编码。
	Code *string `json:"code,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 数据库名。
	DbName *string `json:"db_name,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType *string `json:"dw_type,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须。
	QueueName *string `json:"queue_name,omitempty"`

	// DWS类型需要。
	Schema *string `json:"schema,omitempty"`

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

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	DevPhysicalTable *SyncStatusEnum `json:"dev_physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	DlfTask *SyncStatusEnum `json:"dlf_task,omitempty"`

	Materialization *SyncStatusEnum `json:"materialization,omitempty"`

	PublishToDlm *SyncStatusEnum `json:"publish_to_dlm,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// 标准数量，只读，填写String类型替代Long类型。
	StandardCount *string `json:"standard_count,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 汇总表API ID。
	ApiId *string `json:"api_id,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工作空间名称。
	WorkspaceName *string `json:"workspace_name,omitempty"`

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

func (o AllTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllTableVo struct{}"
	}

	return strings.Join([]string{"AllTableVo", string(data)}, " ")
}
