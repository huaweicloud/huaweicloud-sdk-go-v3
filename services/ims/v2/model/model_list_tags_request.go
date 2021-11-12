package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTagsRequest struct {
	// 用于分页，表示查询几条记录，取值为整数，默认为所有。

	Limit *int32 `json:"limit,omitempty"`
	// 页码，表示需要查询第几页的数据。默认值为1。

	Page *int32 `json:"page,omitempty"`
	// 镜像类型，目前支持以下类型：公共镜像：gold私有镜像：private共享镜像：shared

	Imagetype *ListTagsRequestImagetype `json:"__imagetype,omitempty"`
	// 镜像ID。

	Id *string `json:"id,omitempty"`
	// 镜像状态。取值如下： queued：表示镜像元数据已经创建成功，等待上传镜像文件。 saving：表示镜像正在上传文件到后端存储。 deleted：表示镜像已经删除。 killed：表示镜像上传错误。 active：表示镜像可以正常使用。

	Status *ListTagsRequestStatus `json:"status,omitempty"`
	// 镜像名称。

	Name *string `json:"name,omitempty"`
	// 镜像运行需要的最小磁盘，单位为GB 。

	MinDisk *int32 `json:"min_disk,omitempty"`
	// 镜像平台分类。

	Platform *string `json:"__platform,omitempty"`
	// 镜像系统类型，取值如下：Linux,Windows,Other

	OsType *ListTagsRequestOsType `json:"__os_type,omitempty"`
	// 成员状态。目前取值有accepted、rejected、pending。

	MemberStatus *ListTagsRequestMemberStatus `json:"member_status,omitempty"`
	// 镜像使用环境类型：FusionCompute、Ironic、DataImage。

	VirtualEnvType *ListTagsRequestVirtualEnvType `json:"virtual_env_type,omitempty"`
	// 表示查询某个企业项目下的镜像。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 镜像架构类型。取值包括：x86，arm

	Architecture *ListTagsRequestArchitecture `json:"architecture,omitempty"`
	// 镜像创建时间。支持按照时间点过滤查询，取值格式为“操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询创建时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： created_at=gt:2018-10-28T10:00:00Z

	CreatedAt *string `json:"created_at,omitempty"`
	// 镜像修改时间。支持按照时间点过滤查询，取值格式为“ 操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询修改时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： updated_at=gt:2018-10-28T10:00:00Z

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestImagetype struct {
	value string
}

type ListTagsRequestImagetypeEnum struct {
	GOLD    ListTagsRequestImagetype
	PRIVATE ListTagsRequestImagetype
	SHARED  ListTagsRequestImagetype
}

func GetListTagsRequestImagetypeEnum() ListTagsRequestImagetypeEnum {
	return ListTagsRequestImagetypeEnum{
		GOLD: ListTagsRequestImagetype{
			value: "gold",
		},
		PRIVATE: ListTagsRequestImagetype{
			value: "private",
		},
		SHARED: ListTagsRequestImagetype{
			value: "shared",
		},
	}
}

func (c ListTagsRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestImagetype) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestStatus struct {
	value string
}

type ListTagsRequestStatusEnum struct {
	QUEUED  ListTagsRequestStatus
	SAVING  ListTagsRequestStatus
	DELETED ListTagsRequestStatus
	KILLED  ListTagsRequestStatus
	ACTIVE  ListTagsRequestStatus
}

func GetListTagsRequestStatusEnum() ListTagsRequestStatusEnum {
	return ListTagsRequestStatusEnum{
		QUEUED: ListTagsRequestStatus{
			value: "queued",
		},
		SAVING: ListTagsRequestStatus{
			value: "saving",
		},
		DELETED: ListTagsRequestStatus{
			value: "deleted",
		},
		KILLED: ListTagsRequestStatus{
			value: "killed",
		},
		ACTIVE: ListTagsRequestStatus{
			value: "active",
		},
	}
}

func (c ListTagsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestOsType struct {
	value string
}

type ListTagsRequestOsTypeEnum struct {
	LINUX   ListTagsRequestOsType
	WINDOWS ListTagsRequestOsType
	OTHER   ListTagsRequestOsType
}

func GetListTagsRequestOsTypeEnum() ListTagsRequestOsTypeEnum {
	return ListTagsRequestOsTypeEnum{
		LINUX: ListTagsRequestOsType{
			value: "Linux",
		},
		WINDOWS: ListTagsRequestOsType{
			value: "Windows",
		},
		OTHER: ListTagsRequestOsType{
			value: "Other",
		},
	}
}

func (c ListTagsRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestOsType) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestMemberStatus struct {
	value string
}

type ListTagsRequestMemberStatusEnum struct {
	ACCEPTED ListTagsRequestMemberStatus
	REJECTED ListTagsRequestMemberStatus
	PENDING  ListTagsRequestMemberStatus
}

func GetListTagsRequestMemberStatusEnum() ListTagsRequestMemberStatusEnum {
	return ListTagsRequestMemberStatusEnum{
		ACCEPTED: ListTagsRequestMemberStatus{
			value: "accepted",
		},
		REJECTED: ListTagsRequestMemberStatus{
			value: "rejected",
		},
		PENDING: ListTagsRequestMemberStatus{
			value: "pending",
		},
	}
}

func (c ListTagsRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestMemberStatus) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestVirtualEnvType struct {
	value string
}

type ListTagsRequestVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ListTagsRequestVirtualEnvType
	IRONIC         ListTagsRequestVirtualEnvType
	DATA_IMAGE     ListTagsRequestVirtualEnvType
}

func GetListTagsRequestVirtualEnvTypeEnum() ListTagsRequestVirtualEnvTypeEnum {
	return ListTagsRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListTagsRequestVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ListTagsRequestVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ListTagsRequestVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ListTagsRequestVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestArchitecture struct {
	value string
}

type ListTagsRequestArchitectureEnum struct {
	X86 ListTagsRequestArchitecture
	ARM ListTagsRequestArchitecture
}

func GetListTagsRequestArchitectureEnum() ListTagsRequestArchitectureEnum {
	return ListTagsRequestArchitectureEnum{
		X86: ListTagsRequestArchitecture{
			value: "x86",
		},
		ARM: ListTagsRequestArchitecture{
			value: "arm",
		},
	}
}

func (c ListTagsRequestArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestArchitecture) UnmarshalJSON(b []byte) error {
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
