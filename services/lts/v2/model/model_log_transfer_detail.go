package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogTransferDetail struct {

	// **参数解释：**  OBS转储时间。当创建OBS转储时，必填此参数。与obs_period_unit组合，即\"obs_period\"+\"obs_period_unit\"，必须是\"2min\", \"5min\", \"30min\", \"1hour\", \"3hour\", \"6hour\",\"12hour\"。 **约束限制：**  不涉及。  **取值范围：**  1,2,3,5,6,12,30
	ObsPeriod int32 `json:"obs_period"`

	// **参数解释：**  OBS转储单位。当创建OBS转储时，必填此参数。与obs_period组合，即\"obs_period\"+\"obs_period_unit\"，必须是\"2min\", \"5min\", \"30min\", \"1hour\", \"3hour\", \"6hour\",\"12hour\"。 **约束限制：**  不涉及。  **取值范围：**  min,hour
	ObsPeriodUnit string `json:"obs_period_unit"`

	// **参数解释：**  OBS转储日志桶名称。当创建OBS转储时，必填此参数。 **约束限制：**  不涉及
	ObsBucketName string `json:"obs_bucket_name"`

	// **参数解释：**  OBS转储KMS密钥ID。根据OBS转储日志桶是否加密判断，若OBS转储日志加密桶则必须填写该参数，若OBS转储日志桶则不需要此参数 **约束限制：**  不涉及
	ObsEncryptedId *string `json:"obs_encrypted_id,omitempty"`

	// **参数解释：**  OBS转储自定义转储路径。当创建OBS转储时，根据需要选填此参数。 **约束限制：**  不涉及
	ObsDirPreFixName *string `json:"obs_dir_pre_fix_name,omitempty"`

	// **参数解释：**  OBS转储日志文件前缀。当创建OBS转储时，根据需要选填此参数。 **约束限制：**  不涉及
	ObsPrefixName *string `json:"obs_prefix_name,omitempty"`

	// **参数解释：**  OBS转储时区。参数选择参考OBS转储时区表。如果选择该参数，则必须选择obs_time_zone_id。 **约束限制：**  不涉及
	ObsTimeZone *string `json:"obs_time_zone,omitempty"`

	// **参数解释：**  OBS转储时区ID。参数选择参考OBS转储时区表。如果选择该参数，则必须选择obs_time_zone。 **约束限制：**  不涉及
	ObsTimeZoneId *string `json:"obs_time_zone_id,omitempty"`

	// **参数解释：**  DIS转储通道ID。当创建DIS转储时，必填此参数。 **约束限制：**  不涉及
	DisId *string `json:"dis_id,omitempty"`

	// **参数解释：**  DIS转储通道名称。当创建DIS转储时，必填此参数。 **约束限制：**  不涉及
	DisName *string `json:"dis_name,omitempty"`

	// **参数解释：**  DMS转储kafka ID。当创建DMS转储时，必填此参数。创建DMS转储前，需要使用kafka ID以及kafka Topic进行实例注册。详情见接口注册DMSkafka实例 **约束限制：**  不涉及
	KafkaId *string `json:"kafka_id,omitempty"`

	// **参数解释：**  DMS转储kafka topic。当创建DMS转储时，必填此参数。创建DMS转储前，需要使用kafka ID以及kafka Topic进行实例注册。详情见接口注册DMSkafka实例 **约束限制：**  不涉及
	KafkaTopic *string `json:"kafka_topic,omitempty"`

	// **参数解释：**  OBS企业项目ID。 **约束限制：**  不涉及
	ObsEpsId *string `json:"obs_eps_id,omitempty"`

	// **参数解释：**  OBS是否开启加密。 **约束限制：**  不涉及
	ObsEncryptedEnable *bool `json:"obs_encrypted_enable,omitempty"`

	// **参数解释：**  若开启tag投递，该字段必须包含主机信息：hostIP、hostId、hostName、pathFile、collectTime； 公共字段有：logStreamName、regionName、logGroupName、projectId，为可选填；开启转储标签：streamTag，可选填 **约束限制：**  不涉及
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释：**  dms转储JSON格式选填，可以转储tag字段 **约束限制：**  不涉及
	LtsTags *[]string `json:"lts_tags,omitempty"`

	// **参数解释：**  dms转储JSON格式选填，可以转储日志流标签字段 **约束限制：**  不涉及
	StreamTags *[]string `json:"stream_tags,omitempty"`

	// **参数解释：**  dms转储JSON格式选填，可以转储结构化字段 **约束限制：**  不涉及
	StructFields *[]string `json:"struct_fields,omitempty"`

	// **参数解释：**  dms转储JSON格式选填，无效字段填充 **约束限制：**  不涉及
	InvalidFieldValue *string `json:"invalid_field_value,omitempty"`
}

func (o LogTransferDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogTransferDetail struct{}"
	}

	return strings.Join([]string{"LogTransferDetail", string(data)}, " ")
}
