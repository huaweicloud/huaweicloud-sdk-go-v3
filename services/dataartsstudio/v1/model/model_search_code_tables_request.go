package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchCodeTablesRequest Request Object
type SearchCodeTablesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 目录ID
	DirectoryId *int64 `json:"directory_id,omitempty"`

	// 业务状态
	Status *SearchCodeTablesRequestStatus `json:"status,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchCodeTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCodeTablesRequest struct{}"
	}

	return strings.Join([]string{"SearchCodeTablesRequest", string(data)}, " ")
}

type SearchCodeTablesRequestStatus struct {
	value string
}

type SearchCodeTablesRequestStatusEnum struct {
	DRAFT              SearchCodeTablesRequestStatus
	PUBLISH_DEVELOPING SearchCodeTablesRequestStatus
	PUBLISHED          SearchCodeTablesRequestStatus
	OFFLINE_DEVELOPING SearchCodeTablesRequestStatus
	OFFLINE            SearchCodeTablesRequestStatus
	REJECT             SearchCodeTablesRequestStatus
}

func GetSearchCodeTablesRequestStatusEnum() SearchCodeTablesRequestStatusEnum {
	return SearchCodeTablesRequestStatusEnum{
		DRAFT: SearchCodeTablesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: SearchCodeTablesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: SearchCodeTablesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: SearchCodeTablesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: SearchCodeTablesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: SearchCodeTablesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c SearchCodeTablesRequestStatus) Value() string {
	return c.value
}

func (c SearchCodeTablesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchCodeTablesRequestStatus) UnmarshalJSON(b []byte) error {
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
