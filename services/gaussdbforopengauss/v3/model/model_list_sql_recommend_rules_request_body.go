package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListSqlRecommendRulesRequestBody struct {

	// **参数解释**: 推荐类型。 **约束限制**: 不涉及。 **取值范围**: - all：全部 - exec_count：执行次数 - avg_exec_time：平均执行时间 - max_exec_time：最大执行时间  **默认取值**: all
	RecommendType *ListSqlRecommendRulesRequestBodyRecommendType `json:"recommend_type,omitempty"`

	// **参数解释**: 推荐规则返回条数。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	RecommendCount *int32 `json:"recommend_count,omitempty"`

	// **参数解释**: 是否使用紧急通道。 **约束限制**: 不涉及。 **取值范围**: - true：开启紧急通道 - false：关闭紧急通道  **默认取值**: false
	UseOpsTunnel *bool `json:"use_ops_tunnel,omitempty"`
}

func (o ListSqlRecommendRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlRecommendRulesRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlRecommendRulesRequestBody", string(data)}, " ")
}

type ListSqlRecommendRulesRequestBodyRecommendType struct {
	value string
}

type ListSqlRecommendRulesRequestBodyRecommendTypeEnum struct {
	ALL           ListSqlRecommendRulesRequestBodyRecommendType
	EXEC_COUNT    ListSqlRecommendRulesRequestBodyRecommendType
	AVG_EXEC_TIME ListSqlRecommendRulesRequestBodyRecommendType
	MAX_EXEC_TIME ListSqlRecommendRulesRequestBodyRecommendType
}

func GetListSqlRecommendRulesRequestBodyRecommendTypeEnum() ListSqlRecommendRulesRequestBodyRecommendTypeEnum {
	return ListSqlRecommendRulesRequestBodyRecommendTypeEnum{
		ALL: ListSqlRecommendRulesRequestBodyRecommendType{
			value: "all",
		},
		EXEC_COUNT: ListSqlRecommendRulesRequestBodyRecommendType{
			value: "exec_count",
		},
		AVG_EXEC_TIME: ListSqlRecommendRulesRequestBodyRecommendType{
			value: "avg_exec_time",
		},
		MAX_EXEC_TIME: ListSqlRecommendRulesRequestBodyRecommendType{
			value: "max_exec_time",
		},
	}
}

func (c ListSqlRecommendRulesRequestBodyRecommendType) Value() string {
	return c.value
}

func (c ListSqlRecommendRulesRequestBodyRecommendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlRecommendRulesRequestBodyRecommendType) UnmarshalJSON(b []byte) error {
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
