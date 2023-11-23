package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncTaskInfo 查询同步任务列表时返回的单个同步任务信息实体
type SyncTaskInfo struct {

	// 同步任务ID
	SyncTaskId *string `json:"sync_task_id,omitempty"`

	// 源端云服务提供商。  可选值有AWS、Azure、Aliyun、Tencent、HuaweiCloud、QingCloud、KingsoftCloud、Baidu、Qiniu、UCloud。默认值为Aliyun。
	SrcCloudType *SyncTaskInfoSrcCloudType `json:"src_cloud_type,omitempty"`

	// 源端桶所处的区域
	SrcRegion *string `json:"src_region,omitempty"`

	// 源端桶
	SrcBucket *string `json:"src_bucket,omitempty"`

	// 同步任务创建时间（Unix时间戳，毫秒）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最近启动同步任务时间（Unix时间戳，毫秒）
	LastStartTime *int64 `json:"last_start_time,omitempty"`

	// 目的端桶。
	DstBucket *string `json:"dst_bucket,omitempty"`

	// 目的端region
	DstRegion *string `json:"dst_region,omitempty"`

	// 任务描述，不能超过255个字符，且不能包含<>()\"'&等特殊字符。
	Description *string `json:"description,omitempty"`

	// 同步任务状态 SYNCHRONIZING：同步中 STOPPED：已停止
	Status *SyncTaskInfoStatus `json:"status,omitempty"`

	// 是否开启KMS加密，默认不开启。
	EnableKms *bool `json:"enable_kms,omitempty"`

	// 是否启用元数据迁移，默认否。不启用时，为保证迁移任务正常运行，仍将为您迁移ContentType元数据。
	EnableMetadataMigration *bool `json:"enable_metadata_migration,omitempty"`

	// 是否自动解冻归档数据，默认否。 开启后，如果遇到归档类型数据，会自动解冻再进行迁移。
	EnableRestore *bool `json:"enable_restore,omitempty"`

	// 当源端为腾讯云时，需要填写此参数。
	AppId *string `json:"app_id,omitempty"`

	SourceCdn *SourceCdnResp `json:"source_cdn,omitempty"`

	// 一致性校验方式，用于迁移前/后校验对象是否一致，所有校验方式需满足源端/目的端对象的加密状态一致，具体校验方式和校验结果可通过对象列表查看。默认size_last_modified。 size_last_modified：默认配置。迁移前后，通过对比源端和目的端对象大小+最后修改时间，判断对象是否已存在或迁移后数据是否完整。源端与目的端同名对象大小相同，且目的端对象的最后修改时间不早于源端对象的最后修改时间，则代表该对象已存在/迁移成功。 crc64：目前仅支持华为/阿里/腾讯。迁移前后，通过对比源端和目的端对象元数据中CRC64值是否相同，判断对象是否已存在/迁移完成。如果源端与目的端对象元数据中不存在CRC64值，则系统会默认使用大小/最后修改时间校验方式来校验。 no_check：目前仅支持HTTP/HTTPS数据源。当源端对象无法通过标准http协议中content-length字段获取数据大小时，默认数据下载成功即迁移成功，不对数据做额外校验，且迁移时源端对象默认覆盖目的端同名对象。当源端对象能正常通过标准http协议中content-length字段获取数据大小时，则采用大小/最后修改时间校验方式来校验。
	ConsistencyCheck *SyncTaskInfoConsistencyCheck `json:"consistency_check,omitempty"`
}

func (o SyncTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncTaskInfo struct{}"
	}

	return strings.Join([]string{"SyncTaskInfo", string(data)}, " ")
}

type SyncTaskInfoSrcCloudType struct {
	value string
}

type SyncTaskInfoSrcCloudTypeEnum struct {
	AWS            SyncTaskInfoSrcCloudType
	AZURE          SyncTaskInfoSrcCloudType
	ALIYUN         SyncTaskInfoSrcCloudType
	TENCENT        SyncTaskInfoSrcCloudType
	HUAWEI_CLOUD   SyncTaskInfoSrcCloudType
	QING_CLOUD     SyncTaskInfoSrcCloudType
	KINGSOFT_CLOUD SyncTaskInfoSrcCloudType
	BAIDU          SyncTaskInfoSrcCloudType
	QINIU          SyncTaskInfoSrcCloudType
	U_CLOUD        SyncTaskInfoSrcCloudType
}

func GetSyncTaskInfoSrcCloudTypeEnum() SyncTaskInfoSrcCloudTypeEnum {
	return SyncTaskInfoSrcCloudTypeEnum{
		AWS: SyncTaskInfoSrcCloudType{
			value: "AWS",
		},
		AZURE: SyncTaskInfoSrcCloudType{
			value: "Azure",
		},
		ALIYUN: SyncTaskInfoSrcCloudType{
			value: "Aliyun",
		},
		TENCENT: SyncTaskInfoSrcCloudType{
			value: "Tencent",
		},
		HUAWEI_CLOUD: SyncTaskInfoSrcCloudType{
			value: "HuaweiCloud",
		},
		QING_CLOUD: SyncTaskInfoSrcCloudType{
			value: "QingCloud",
		},
		KINGSOFT_CLOUD: SyncTaskInfoSrcCloudType{
			value: "KingsoftCloud",
		},
		BAIDU: SyncTaskInfoSrcCloudType{
			value: "Baidu",
		},
		QINIU: SyncTaskInfoSrcCloudType{
			value: "Qiniu",
		},
		U_CLOUD: SyncTaskInfoSrcCloudType{
			value: "UCloud",
		},
	}
}

func (c SyncTaskInfoSrcCloudType) Value() string {
	return c.value
}

func (c SyncTaskInfoSrcCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncTaskInfoSrcCloudType) UnmarshalJSON(b []byte) error {
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

type SyncTaskInfoStatus struct {
	value string
}

type SyncTaskInfoStatusEnum struct {
	SYNCHRONIZING SyncTaskInfoStatus
	STOPPED       SyncTaskInfoStatus
}

func GetSyncTaskInfoStatusEnum() SyncTaskInfoStatusEnum {
	return SyncTaskInfoStatusEnum{
		SYNCHRONIZING: SyncTaskInfoStatus{
			value: "SYNCHRONIZING",
		},
		STOPPED: SyncTaskInfoStatus{
			value: "STOPPED",
		},
	}
}

func (c SyncTaskInfoStatus) Value() string {
	return c.value
}

func (c SyncTaskInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncTaskInfoStatus) UnmarshalJSON(b []byte) error {
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

type SyncTaskInfoConsistencyCheck struct {
	value string
}

type SyncTaskInfoConsistencyCheckEnum struct {
	SIZE_LAST_MODIFIED SyncTaskInfoConsistencyCheck
	CRC64              SyncTaskInfoConsistencyCheck
	NO_CHECK           SyncTaskInfoConsistencyCheck
}

func GetSyncTaskInfoConsistencyCheckEnum() SyncTaskInfoConsistencyCheckEnum {
	return SyncTaskInfoConsistencyCheckEnum{
		SIZE_LAST_MODIFIED: SyncTaskInfoConsistencyCheck{
			value: "size_last_modified",
		},
		CRC64: SyncTaskInfoConsistencyCheck{
			value: "crc64",
		},
		NO_CHECK: SyncTaskInfoConsistencyCheck{
			value: "no_check",
		},
	}
}

func (c SyncTaskInfoConsistencyCheck) Value() string {
	return c.value
}

func (c SyncTaskInfoConsistencyCheck) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncTaskInfoConsistencyCheck) UnmarshalJSON(b []byte) error {
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
