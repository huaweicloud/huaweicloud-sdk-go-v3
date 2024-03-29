package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchSubjectRequest Request Object
type SearchSubjectRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按负责人查询
	Owner *string `json:"owner,omitempty"`

	// 业务状态
	Status *SearchSubjectRequestStatus `json:"status,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 父目录ID，根节点没有此ID，空值为所有，-1为根节点下节点
	ParentId *int64 `json:"parent_id,omitempty"`
}

func (o SearchSubjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectRequest struct{}"
	}

	return strings.Join([]string{"SearchSubjectRequest", string(data)}, " ")
}

type SearchSubjectRequestStatus struct {
	value string
}

type SearchSubjectRequestStatusEnum struct {
	DRAFT              SearchSubjectRequestStatus
	PUBLISH_DEVELOPING SearchSubjectRequestStatus
	PUBLISHED          SearchSubjectRequestStatus
	OFFLINE_DEVELOPING SearchSubjectRequestStatus
	OFFLINE            SearchSubjectRequestStatus
	REJECT             SearchSubjectRequestStatus
}

func GetSearchSubjectRequestStatusEnum() SearchSubjectRequestStatusEnum {
	return SearchSubjectRequestStatusEnum{
		DRAFT: SearchSubjectRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: SearchSubjectRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: SearchSubjectRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: SearchSubjectRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: SearchSubjectRequestStatus{
			value: "OFFLINE",
		},
		REJECT: SearchSubjectRequestStatus{
			value: "REJECT",
		},
	}
}

func (c SearchSubjectRequestStatus) Value() string {
	return c.value
}

func (c SearchSubjectRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchSubjectRequestStatus) UnmarshalJSON(b []byte) error {
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
