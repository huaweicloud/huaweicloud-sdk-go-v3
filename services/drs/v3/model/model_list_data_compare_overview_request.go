package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDataCompareOverviewRequest Request Object
type ListDataCompareOverviewRequest struct {

	// 请求语言类型。
	XLanguage *ListDataCompareOverviewRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`

	// 对比状态。 - 0：对比不一致 - 2：对比一致 - 3：目标库表不存在 - 4：对比失败 - 5：正在对比中 - 6：等待对比中 - 7：任务已取消 - 8：源库为空 - 9：目标库为空 - 10：源库和目标库都为空 - 11：源表不存在 - 12：目标表不存在 - 13：原表和目标表都不存在 - 14：源数据库连接失败 - 15：目标库数据库连接失败 - 16：源数据库执行SQL超时 - 17：目标数据库执行SQL超时 - 18：源数据库执行SQL错误 - 19：目标数据库执行SQL错误 - 20：源库和目标库都不存在 - 21：源库不存在 - 22：目标库不存在 - 23：行数为亿行，未进行对比 - 27：超时
	Status *int32 `json:"status,omitempty"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDataCompareOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataCompareOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListDataCompareOverviewRequest", string(data)}, " ")
}

type ListDataCompareOverviewRequestXLanguage struct {
	value string
}

type ListDataCompareOverviewRequestXLanguageEnum struct {
	EN_US ListDataCompareOverviewRequestXLanguage
	ZH_CN ListDataCompareOverviewRequestXLanguage
}

func GetListDataCompareOverviewRequestXLanguageEnum() ListDataCompareOverviewRequestXLanguageEnum {
	return ListDataCompareOverviewRequestXLanguageEnum{
		EN_US: ListDataCompareOverviewRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListDataCompareOverviewRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListDataCompareOverviewRequestXLanguage) Value() string {
	return c.value
}

func (c ListDataCompareOverviewRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDataCompareOverviewRequestXLanguage) UnmarshalJSON(b []byte) error {
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
