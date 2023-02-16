package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AskQuestionReq struct {

	// 用户输入问题
	Question string `json:"question"`

	// 主题列表
	Domains *[]string `json:"domains,omitempty"`

	// 返回匹配度最高的数据条数
	Top *int32 `json:"top,omitempty"`

	// 会话ID
	SessionId string `json:"session_id"`

	// 语料ID
	SourceQaPairId *string `json:"source_qa_pair_id,omitempty"`

	// 操作类型:0-手动输入，1-单击热点问题，2-单击猜你想问，3-单击推荐问题，4-单击问题提示
	OperateType *AskQuestionReqOperateType `json:"operate_type,omitempty"`

	// 是否启用内部阈值开关
	ThresholdEnable *string `json:"threshold_enable,omitempty"`

	// 产品类型id
	ProductTypeId *string `json:"product_type_id,omitempty"`
}

func (o AskQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AskQuestionReq struct{}"
	}

	return strings.Join([]string{"AskQuestionReq", string(data)}, " ")
}

type AskQuestionReqOperateType struct {
	value int32
}

type AskQuestionReqOperateTypeEnum struct {
	E_0 AskQuestionReqOperateType
	E_1 AskQuestionReqOperateType
	E_2 AskQuestionReqOperateType
	E_3 AskQuestionReqOperateType
	E_4 AskQuestionReqOperateType
}

func GetAskQuestionReqOperateTypeEnum() AskQuestionReqOperateTypeEnum {
	return AskQuestionReqOperateTypeEnum{
		E_0: AskQuestionReqOperateType{
			value: 0,
		}, E_1: AskQuestionReqOperateType{
			value: 1,
		}, E_2: AskQuestionReqOperateType{
			value: 2,
		}, E_3: AskQuestionReqOperateType{
			value: 3,
		}, E_4: AskQuestionReqOperateType{
			value: 4,
		},
	}
}

func (c AskQuestionReqOperateType) Value() int32 {
	return c.value
}

func (c AskQuestionReqOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AskQuestionReqOperateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
