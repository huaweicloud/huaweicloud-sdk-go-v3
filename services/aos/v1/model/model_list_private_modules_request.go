package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateModulesRequest Request Object
type ListPrivateModulesRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListPrivateModulesRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListPrivateModulesRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListPrivateModulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateModulesRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateModulesRequest", string(data)}, " ")
}

type ListPrivateModulesRequestSortKey struct {
	value string
}

type ListPrivateModulesRequestSortKeyEnum struct {
	CREATE_TIME ListPrivateModulesRequestSortKey
}

func GetListPrivateModulesRequestSortKeyEnum() ListPrivateModulesRequestSortKeyEnum {
	return ListPrivateModulesRequestSortKeyEnum{
		CREATE_TIME: ListPrivateModulesRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListPrivateModulesRequestSortKey) Value() string {
	return c.value
}

func (c ListPrivateModulesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateModulesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListPrivateModulesRequestSortDir struct {
	value string
}

type ListPrivateModulesRequestSortDirEnum struct {
	ASC  ListPrivateModulesRequestSortDir
	DESC ListPrivateModulesRequestSortDir
}

func GetListPrivateModulesRequestSortDirEnum() ListPrivateModulesRequestSortDirEnum {
	return ListPrivateModulesRequestSortDirEnum{
		ASC: ListPrivateModulesRequestSortDir{
			value: "asc",
		},
		DESC: ListPrivateModulesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPrivateModulesRequestSortDir) Value() string {
	return c.value
}

func (c ListPrivateModulesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateModulesRequestSortDir) UnmarshalJSON(b []byte) error {
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
