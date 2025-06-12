package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateProviderVersionsRequest Request Object
type ListPrivateProviderVersionsRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有provider（private-provider）的名称。此名字在domain_id+region下应唯一，可以使用小写英文、数字、中划线。仅支持以小写英文、数字开头结尾。
	ProviderName string `json:"provider_name"`

	// 私有provider的唯一Id，由资源编排服务随机生成，为UUID。  由于私有provider名称仅仅在同一时间下唯一，即用户允许先生成一个叫helloword的私有provider，删除，再重新创建一个同名私有provider。  对于团队并行开发，用户可能希望确保，当前我操作的私有provider就是我以为的那个，而不是由其他队友删除后创建的同名私有provider。  因此，使用ID就可以做到强匹配。资源编排服务保证每次创建私有provider所对应的Id都不相同。  如果给予的provider_id和当前私有provider的Id不一致，则返回400。
	ProviderId *string `json:"provider_id,omitempty"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListPrivateProviderVersionsRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListPrivateProviderVersionsRequestSortDir `json:"sort_dir,omitempty"`

	// 分页标记。当一页无法返回所有结果，上一次的请求将返回next_marker以指引还有更多页数，用户可以将next_marker中的值放到此处以查询下一页的信息。此marker只能用于与上一请求指定的相同参数的请求。不指定时默认从第一页开始查询。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的最多结果数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPrivateProviderVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateProviderVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateProviderVersionsRequest", string(data)}, " ")
}

type ListPrivateProviderVersionsRequestSortKey struct {
	value string
}

type ListPrivateProviderVersionsRequestSortKeyEnum struct {
	CREATE_TIME ListPrivateProviderVersionsRequestSortKey
}

func GetListPrivateProviderVersionsRequestSortKeyEnum() ListPrivateProviderVersionsRequestSortKeyEnum {
	return ListPrivateProviderVersionsRequestSortKeyEnum{
		CREATE_TIME: ListPrivateProviderVersionsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListPrivateProviderVersionsRequestSortKey) Value() string {
	return c.value
}

func (c ListPrivateProviderVersionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateProviderVersionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListPrivateProviderVersionsRequestSortDir struct {
	value string
}

type ListPrivateProviderVersionsRequestSortDirEnum struct {
	ASC  ListPrivateProviderVersionsRequestSortDir
	DESC ListPrivateProviderVersionsRequestSortDir
}

func GetListPrivateProviderVersionsRequestSortDirEnum() ListPrivateProviderVersionsRequestSortDirEnum {
	return ListPrivateProviderVersionsRequestSortDirEnum{
		ASC: ListPrivateProviderVersionsRequestSortDir{
			value: "asc",
		},
		DESC: ListPrivateProviderVersionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPrivateProviderVersionsRequestSortDir) Value() string {
	return c.value
}

func (c ListPrivateProviderVersionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateProviderVersionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
