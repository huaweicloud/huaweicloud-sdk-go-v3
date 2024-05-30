package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApproverVo struct {

	// 审批单ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 审批人姓名。
	ApproverName *string `json:"approver_name,omitempty"`

	// 审批人ID。
	UserId *string `json:"user_id,omitempty"`

	// 审批人名称。
	UserName *string `json:"user_name,omitempty"`

	// email信息。
	Email *string `json:"email,omitempty"`

	// 用户类型。 枚举值：   - BIZ_METRIC_OWNER: 业务指标责任人   - APPROVER: 审批人   - BIZ_METRIC_OWNER_AND_APPROVER: 业务指标责任人是审核人
	UserType *ApproverVoUserType `json:"user_type,omitempty"`

	// 电话号码。
	PhoneNumber *string `json:"phone_number,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 业务系统名称。
	AppName *string `json:"app_name,omitempty"`

	// smn主题urn。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ApproverVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverVo struct{}"
	}

	return strings.Join([]string{"ApproverVo", string(data)}, " ")
}

type ApproverVoUserType struct {
	value string
}

type ApproverVoUserTypeEnum struct {
	BIZ_METRIC_OWNER              ApproverVoUserType
	APPROVER                      ApproverVoUserType
	BIZ_METRIC_OWNER_AND_APPROVER ApproverVoUserType
}

func GetApproverVoUserTypeEnum() ApproverVoUserTypeEnum {
	return ApproverVoUserTypeEnum{
		BIZ_METRIC_OWNER: ApproverVoUserType{
			value: "BIZ_METRIC_OWNER",
		},
		APPROVER: ApproverVoUserType{
			value: "APPROVER",
		},
		BIZ_METRIC_OWNER_AND_APPROVER: ApproverVoUserType{
			value: "BIZ_METRIC_OWNER_AND_APPROVER",
		},
	}
}

func (c ApproverVoUserType) Value() string {
	return c.value
}

func (c ApproverVoUserType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApproverVoUserType) UnmarshalJSON(b []byte) error {
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
