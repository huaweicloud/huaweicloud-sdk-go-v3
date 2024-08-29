package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListComponentsRequest Request Object
type ListComponentsRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	// 企业项目ID。  - 创建环境时，环境会绑定企业项目ID。      - 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。     - 该字段不传（或传为字符串“0”）时，则查询默认企业项目下的资源。  > 关于企业项目ID的获取及企业项目特性的详细信息，请参见《[企业管理服务用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)》。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 环境ID。      - 获取环境ID，通过《[云应用引擎API参考](https://support.huaweicloud.com/api-cae/ListEnvironments.html)》的“获取环境列表”章节获取环境信息。     - 请求响应成功后在响应体的items数组中的一个元素即为一个环境的信息，其中id字段即是环境ID。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 限制本次返回结果的条数。
	Limit *string `json:"limit,omitempty"`

	// 分页偏移位，查询起始位置。
	Offset *string `json:"offset,omitempty"`

	// 排序字段。
	SortKey *ListComponentsRequestSortKey `json:"sort_key,omitempty"`

	// 升降序规则。
	Sort *ListComponentsRequestSort `json:"sort,omitempty"`
}

func (o ListComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentsRequest struct{}"
	}

	return strings.Join([]string{"ListComponentsRequest", string(data)}, " ")
}

type ListComponentsRequestSortKey struct {
	value string
}

type ListComponentsRequestSortKeyEnum struct {
	CREATED_AT ListComponentsRequestSortKey
	UPDATED_AT ListComponentsRequestSortKey
}

func GetListComponentsRequestSortKeyEnum() ListComponentsRequestSortKeyEnum {
	return ListComponentsRequestSortKeyEnum{
		CREATED_AT: ListComponentsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListComponentsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListComponentsRequestSortKey) Value() string {
	return c.value
}

func (c ListComponentsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListComponentsRequestSort struct {
	value string
}

type ListComponentsRequestSortEnum struct {
	ASC  ListComponentsRequestSort
	DESC ListComponentsRequestSort
}

func GetListComponentsRequestSortEnum() ListComponentsRequestSortEnum {
	return ListComponentsRequestSortEnum{
		ASC: ListComponentsRequestSort{
			value: "asc",
		},
		DESC: ListComponentsRequestSort{
			value: "desc",
		},
	}
}

func (c ListComponentsRequestSort) Value() string {
	return c.value
}

func (c ListComponentsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentsRequestSort) UnmarshalJSON(b []byte) error {
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
