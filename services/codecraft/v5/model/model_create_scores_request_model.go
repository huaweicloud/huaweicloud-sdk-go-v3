package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateScoresRequestModel struct {

	// 大赛ID，大赛平台提供
	CompetitionId string `json:"competition_id"`

	// 大赛阶段ID，大赛平台提供
	StageId string `json:"stage_id"`

	// 第三方服务作品ID
	WorksId int32 `json:"works_id"`

	// 作品名称，名称最大字符数为75，并且不能有含有特殊符号
	Name string `json:"name"`

	// 作品类型,例如docx、png、zip等
	WorksKind *string `json:"works_kind,omitempty"`

	// 作品分数，作品状态为failed时传-1，计算长度时包括小数点，小数点后面最多保留四位
	Score float64 `json:"score"`

	// 作品状态success|failed。判题时，需要对上传作品进行检查，当作品不符合要求时，应该返回failed，并将提示信息通过 message显示出来
	Status CreateScoresRequestModelStatus `json:"status"`

	// 作品创建时间
	CreatedTime string `json:"created_time"`

	// 作品备注信息
	Note *string `json:"note,omitempty"`

	// 作品描述信息
	Message *string `json:"message,omitempty"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`
}

func (o CreateScoresRequestModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScoresRequestModel struct{}"
	}

	return strings.Join([]string{"CreateScoresRequestModel", string(data)}, " ")
}

type CreateScoresRequestModelStatus struct {
	value string
}

type CreateScoresRequestModelStatusEnum struct {
	SUCCESS CreateScoresRequestModelStatus
	FAILED  CreateScoresRequestModelStatus
}

func GetCreateScoresRequestModelStatusEnum() CreateScoresRequestModelStatusEnum {
	return CreateScoresRequestModelStatusEnum{
		SUCCESS: CreateScoresRequestModelStatus{
			value: "success",
		},
		FAILED: CreateScoresRequestModelStatus{
			value: "failed",
		},
	}
}

func (c CreateScoresRequestModelStatus) Value() string {
	return c.value
}

func (c CreateScoresRequestModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScoresRequestModelStatus) UnmarshalJSON(b []byte) error {
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
