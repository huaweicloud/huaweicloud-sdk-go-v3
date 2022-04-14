package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBackupsRequest struct {
	// 还原点ID

	CheckpointId *string `json:"checkpoint_id,omitempty"`
	// 专属云

	Dec *bool `json:"dec,omitempty"`
	// 备份产生时间范围的结束时间，格式为%YYYY-%mm-%ddT%HH:%MM:%SSZ，例如2018-02-01T12:00:00Z

	EndTime *string `json:"end_time,omitempty"`
	// 备份类型

	ImageType *ListBackupsRequestImageType `json:"image_type,omitempty"`
	// 每页显示的条目数量，正整数

	Limit *int32 `json:"limit,omitempty"`
	// 上一次查询最后一条的id

	Marker *string `json:"marker,omitempty"`
	// 名称

	Name *string `json:"name,omitempty"`
	// 偏移值，正整数

	Offset *int32 `json:"offset,omitempty"`
	// 支持按az来过滤

	ResourceAz *string `json:"resource_az,omitempty"`
	// 资源ID

	ResourceId *string `json:"resource_id,omitempty"`
	// 资源名称

	ResourceName *string `json:"resource_name,omitempty"`
	// 资源类型

	ResourceType *ListBackupsRequestResourceType `json:"resource_type,omitempty"`
	// sort的内容为一组由逗号分隔的属性及可选排序方向组成，形如<key1>[:<direction>],<key2>[:<direction>],其中direction的取值为asc (升序) 或 desc (降序),如没有传入direction参数，默认为降序，sort内容的长度限制为255个字符。key取值范围:[created_at，updated_at，name，status，protected_at，id]

	Sort *string `json:"sort,omitempty"`
	// 备份产生时间范围的开始时间，格式为%YYYY-%mm-%ddT%HH:%MM:%SSZ，例如2018-02-01T12:00:00Z

	StartTime *string `json:"start_time,omitempty"`
	// 状态。 调用API时，支持通过传多个status值进行过滤。例如：status=available&status=error

	Status *ListBackupsRequestStatus `json:"status,omitempty"`
	// 存储库ID

	VaultId *string `json:"vault_id,omitempty"`
	// 企业项目id或all_granted_eps，all_granted_eps表示查询用户有权限的所有企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 持有类型，私有的private/共享的shared/全部all_granted，默认只查询private。

	OwnType *ListBackupsRequestOwnType `json:"own_type,omitempty"`
	// 共享状态

	MemberStatus *ListBackupsRequestMemberStatus `json:"member_status,omitempty"`
	// 父备份ID

	ParentId *string `json:"parent_id,omitempty"`
	// 根据存储库使用率过滤备份，取值范围 [1, 100]，含1和100。例如，used_percent=80，表示筛选所属存储库使用率大于等于80%的所有备份。

	UsedPercent *string `json:"used_percent,omitempty"`
	// 是否返回复制记录

	ShowReplication *bool `json:"show_replication,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

type ListBackupsRequestImageType struct {
	value string
}

type ListBackupsRequestImageTypeEnum struct {
	BACKUP      ListBackupsRequestImageType
	REPLICATION ListBackupsRequestImageType
}

func GetListBackupsRequestImageTypeEnum() ListBackupsRequestImageTypeEnum {
	return ListBackupsRequestImageTypeEnum{
		BACKUP: ListBackupsRequestImageType{
			value: "backup",
		},
		REPLICATION: ListBackupsRequestImageType{
			value: "replication",
		},
	}
}

func (c ListBackupsRequestImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestImageType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListBackupsRequestResourceType struct {
	value string
}

type ListBackupsRequestResourceTypeEnum struct {
	OSCINDERVOLUME ListBackupsRequestResourceType
	OSNOVASERVER   ListBackupsRequestResourceType
}

func GetListBackupsRequestResourceTypeEnum() ListBackupsRequestResourceTypeEnum {
	return ListBackupsRequestResourceTypeEnum{
		OSCINDERVOLUME: ListBackupsRequestResourceType{
			value: "OS::Cinder::Volume",
		},
		OSNOVASERVER: ListBackupsRequestResourceType{
			value: "OS::Nova::Server",
		},
	}
}

func (c ListBackupsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListBackupsRequestStatus struct {
	value string
}

type ListBackupsRequestStatusEnum struct {
	AVAILABLE       ListBackupsRequestStatus
	PROTECTING      ListBackupsRequestStatus
	DELETING        ListBackupsRequestStatus
	RESTORING       ListBackupsRequestStatus
	ERROR           ListBackupsRequestStatus
	WAITING_PROTECT ListBackupsRequestStatus
	WAITING_DELETE  ListBackupsRequestStatus
	WAITING_RESTORE ListBackupsRequestStatus
}

func GetListBackupsRequestStatusEnum() ListBackupsRequestStatusEnum {
	return ListBackupsRequestStatusEnum{
		AVAILABLE: ListBackupsRequestStatus{
			value: "available",
		},
		PROTECTING: ListBackupsRequestStatus{
			value: "protecting",
		},
		DELETING: ListBackupsRequestStatus{
			value: "deleting",
		},
		RESTORING: ListBackupsRequestStatus{
			value: "restoring",
		},
		ERROR: ListBackupsRequestStatus{
			value: "error",
		},
		WAITING_PROTECT: ListBackupsRequestStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: ListBackupsRequestStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: ListBackupsRequestStatus{
			value: "waiting_restore",
		},
	}
}

func (c ListBackupsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListBackupsRequestOwnType struct {
	value string
}

type ListBackupsRequestOwnTypeEnum struct {
	ALL_GRANTED ListBackupsRequestOwnType
	PRIVATE     ListBackupsRequestOwnType
	SHARED      ListBackupsRequestOwnType
}

func GetListBackupsRequestOwnTypeEnum() ListBackupsRequestOwnTypeEnum {
	return ListBackupsRequestOwnTypeEnum{
		ALL_GRANTED: ListBackupsRequestOwnType{
			value: "all_granted",
		},
		PRIVATE: ListBackupsRequestOwnType{
			value: "private",
		},
		SHARED: ListBackupsRequestOwnType{
			value: "shared",
		},
	}
}

func (c ListBackupsRequestOwnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestOwnType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListBackupsRequestMemberStatus struct {
	value string
}

type ListBackupsRequestMemberStatusEnum struct {
	PENDING  ListBackupsRequestMemberStatus
	ACCEPTED ListBackupsRequestMemberStatus
	REJECTED ListBackupsRequestMemberStatus
}

func GetListBackupsRequestMemberStatusEnum() ListBackupsRequestMemberStatusEnum {
	return ListBackupsRequestMemberStatusEnum{
		PENDING: ListBackupsRequestMemberStatus{
			value: "pending",
		},
		ACCEPTED: ListBackupsRequestMemberStatus{
			value: "accepted",
		},
		REJECTED: ListBackupsRequestMemberStatus{
			value: "rejected",
		},
	}
}

func (c ListBackupsRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestMemberStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
