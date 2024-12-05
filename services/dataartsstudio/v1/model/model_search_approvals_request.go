package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchApprovalsRequest Request Object
type SearchApprovalsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 业务定义ID，ID字符串。
	BizId *string `json:"biz_id,omitempty"`

	// 按名称或编码模糊查询。
	Name *string `json:"name,omitempty"`

	// 按创建者查询。
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询。
	Approver *string `json:"approver,omitempty"`

	// 审批单状态。 枚举值：   - DEVELOPING: 待审批   - FINISHED: 已审批
	ApprovalStatus *SearchApprovalsRequestApprovalStatus `json:"approval_status,omitempty"`

	// 审批单状态。 枚举值：   - DEVELOPING: 待审批   - APPROVED: 审批通过   - REJECT: 审批驳回
	ApprovalStatusDetail *SearchApprovalsRequestApprovalStatusDetail `json:"approval_status_detail,omitempty"`

	// 业务审核类型。 枚举值：   - PUBLISH: 发布   - OFFLINE: 下线
	ApprovalType *SearchApprovalsRequestApprovalType `json:"approval_type,omitempty"`

	// 按业务类型查询，可选业务类型有：ATOMIC_INDEX（原子指标）、DERIVATIVE_INDEX（衍生指标）、DIMENSION（维度）、TIME_CONDITION（时间限定）、DIMENSION_LOGIC_TABLE（维度表）、FACT_LOGIC_TABLE（事实表）、AGGREGATION_LOGIC_TABLE（汇总表）、TABLE_MODEL（关系建模表）、CODE_TABLE（码表）、STANDARD_ELEMENT）（数据标准）、BIZ_METRIC（业务指标）、COMPOUND_METRIC（复合指标）、SUBJECT（主题）、ATOMIC_METRIC（原子指标（新））、DERIVED_METRIC（衍生指标（新））、COMPOSITE_METRIC（复合指标（新））。
	BizType *string `json:"biz_type,omitempty"`

	// 时间过滤左边界，与end_time一起使用，只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界，与begin_time一起使用只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchApprovalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApprovalsRequest struct{}"
	}

	return strings.Join([]string{"SearchApprovalsRequest", string(data)}, " ")
}

type SearchApprovalsRequestApprovalStatus struct {
	value string
}

type SearchApprovalsRequestApprovalStatusEnum struct {
	DEVELOPING SearchApprovalsRequestApprovalStatus
	FINISHED   SearchApprovalsRequestApprovalStatus
}

func GetSearchApprovalsRequestApprovalStatusEnum() SearchApprovalsRequestApprovalStatusEnum {
	return SearchApprovalsRequestApprovalStatusEnum{
		DEVELOPING: SearchApprovalsRequestApprovalStatus{
			value: "DEVELOPING",
		},
		FINISHED: SearchApprovalsRequestApprovalStatus{
			value: "FINISHED",
		},
	}
}

func (c SearchApprovalsRequestApprovalStatus) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalStatus) UnmarshalJSON(b []byte) error {
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

type SearchApprovalsRequestApprovalStatusDetail struct {
	value string
}

type SearchApprovalsRequestApprovalStatusDetailEnum struct {
	DEVELOPING SearchApprovalsRequestApprovalStatusDetail
	APPROVED   SearchApprovalsRequestApprovalStatusDetail
	REJECT     SearchApprovalsRequestApprovalStatusDetail
}

func GetSearchApprovalsRequestApprovalStatusDetailEnum() SearchApprovalsRequestApprovalStatusDetailEnum {
	return SearchApprovalsRequestApprovalStatusDetailEnum{
		DEVELOPING: SearchApprovalsRequestApprovalStatusDetail{
			value: "DEVELOPING",
		},
		APPROVED: SearchApprovalsRequestApprovalStatusDetail{
			value: "APPROVED",
		},
		REJECT: SearchApprovalsRequestApprovalStatusDetail{
			value: "REJECT",
		},
	}
}

func (c SearchApprovalsRequestApprovalStatusDetail) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalStatusDetail) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalStatusDetail) UnmarshalJSON(b []byte) error {
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

type SearchApprovalsRequestApprovalType struct {
	value string
}

type SearchApprovalsRequestApprovalTypeEnum struct {
	PUBLISH SearchApprovalsRequestApprovalType
	OFFLINE SearchApprovalsRequestApprovalType
}

func GetSearchApprovalsRequestApprovalTypeEnum() SearchApprovalsRequestApprovalTypeEnum {
	return SearchApprovalsRequestApprovalTypeEnum{
		PUBLISH: SearchApprovalsRequestApprovalType{
			value: "PUBLISH",
		},
		OFFLINE: SearchApprovalsRequestApprovalType{
			value: "OFFLINE",
		},
	}
}

func (c SearchApprovalsRequestApprovalType) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalType) UnmarshalJSON(b []byte) error {
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
