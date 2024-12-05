package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchAtomicIndexesRequest Request Object
type SearchAtomicIndexesRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 按名称或编码模糊查询。
	Name *string `json:"name,omitempty"`

	// 按创建者查询。
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询。
	Approver *string `json:"approver,omitempty"`

	// 业务状态。 枚举值：   - DRAFT: 草稿   - PUBLISH_DEVELOPING: 发布待审批   - PUBLISHED: 已发布   - OFFLINE_DEVELOPING: 下线待审批   - OFFLINE: 已下线   - REJECT: 已驳回
	Status *SearchAtomicIndexesRequestStatus `json:"status,omitempty"`

	// 时间过滤左边界，与end_time一起使用，只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界，与begin_time一起使用只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 业务对象l3的ID，ID字符串。
	L3Id *string `json:"l3_id,omitempty"`

	// 关联表的ID，ID字符串。
	TableId *string `json:"table_id,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchAtomicIndexesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAtomicIndexesRequest struct{}"
	}

	return strings.Join([]string{"SearchAtomicIndexesRequest", string(data)}, " ")
}

type SearchAtomicIndexesRequestStatus struct {
	value string
}

type SearchAtomicIndexesRequestStatusEnum struct {
	DRAFT              SearchAtomicIndexesRequestStatus
	PUBLISH_DEVELOPING SearchAtomicIndexesRequestStatus
	PUBLISHED          SearchAtomicIndexesRequestStatus
	OFFLINE_DEVELOPING SearchAtomicIndexesRequestStatus
	OFFLINE            SearchAtomicIndexesRequestStatus
	REJECT             SearchAtomicIndexesRequestStatus
}

func GetSearchAtomicIndexesRequestStatusEnum() SearchAtomicIndexesRequestStatusEnum {
	return SearchAtomicIndexesRequestStatusEnum{
		DRAFT: SearchAtomicIndexesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: SearchAtomicIndexesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: SearchAtomicIndexesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: SearchAtomicIndexesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: SearchAtomicIndexesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: SearchAtomicIndexesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c SearchAtomicIndexesRequestStatus) Value() string {
	return c.value
}

func (c SearchAtomicIndexesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchAtomicIndexesRequestStatus) UnmarshalJSON(b []byte) error {
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
