package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishVersionVo 版本信息。
type PublishVersionVo struct {

	// 版本ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 版本名称。
	VersionName string `json:"version_name"`

	// 版本标记，只读。
	VersionTag *string `json:"version_tag,omitempty"`

	// 版本描述。
	Description *string `json:"description,omitempty"`

	// 业务对象ID，填写String类型替代Long类型。
	BizId *string `json:"biz_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 业务详情，只读。
	BizInfo *string `json:"biz_info,omitempty"`

	// 业务对象。
	BizInfoVo *interface{} `json:"biz_info_vo,omitempty"`

	// 影响信息，只读。
	EffectObjs *string `json:"effect_objs,omitempty"`

	// 变化信息，只读。
	ChangeProps *string `json:"change_props,omitempty"`

	// SQL脚本，只读。
	SqlDdl *string `json:"sql_ddl,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	DevPhysicalTable *SyncStatusEnum `json:"dev_physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	DlfTask *SyncStatusEnum `json:"dlf_task,omitempty"`

	Materialization *SyncStatusEnum `json:"materialization,omitempty"`

	PublishToDlm *SyncStatusEnum `json:"publish_to_dlm,omitempty"`

	BizMetric *SyncStatusEnum `json:"biz_metric,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// 是否为当前版本，只读。
	IsCurrentVersion *bool `json:"is_current_version,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`
}

func (o PublishVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishVersionVo struct{}"
	}

	return strings.Join([]string{"PublishVersionVo", string(data)}, " ")
}
