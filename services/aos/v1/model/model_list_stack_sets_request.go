package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListStackSetsRequest Request Object
type ListStackSetsRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 过滤条件  * 与（AND）运算符使用逗号（，）定义 * 或（OR）运算符使用竖线（|）定义，OR运算符优先级高于AND运算符 * 不支持括号 * 过滤运算符仅支持双等号（==） * 过滤参数名及其值仅支持包含大小写英文、数字和下划线 * 过滤条件中禁止使用分号，如果有分号，则此条过滤会被忽略 * 一个过滤参数仅能与一个与条件相关，一个与条件中的多个或条件仅能与一个过滤参数相关
	Filter *string `json:"filter,omitempty"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListStackSetsRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListStackSetsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListStackSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackSetsRequest struct{}"
	}

	return strings.Join([]string{"ListStackSetsRequest", string(data)}, " ")
}

type ListStackSetsRequestSortKey struct {
	value string
}

type ListStackSetsRequestSortKeyEnum struct {
	CREATE_TIME ListStackSetsRequestSortKey
}

func GetListStackSetsRequestSortKeyEnum() ListStackSetsRequestSortKeyEnum {
	return ListStackSetsRequestSortKeyEnum{
		CREATE_TIME: ListStackSetsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListStackSetsRequestSortKey) Value() string {
	return c.value
}

func (c ListStackSetsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStackSetsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListStackSetsRequestSortDir struct {
	value string
}

type ListStackSetsRequestSortDirEnum struct {
	ASC  ListStackSetsRequestSortDir
	DESC ListStackSetsRequestSortDir
}

func GetListStackSetsRequestSortDirEnum() ListStackSetsRequestSortDirEnum {
	return ListStackSetsRequestSortDirEnum{
		ASC: ListStackSetsRequestSortDir{
			value: "asc",
		},
		DESC: ListStackSetsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListStackSetsRequestSortDir) Value() string {
	return c.value
}

func (c ListStackSetsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStackSetsRequestSortDir) UnmarshalJSON(b []byte) error {
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
