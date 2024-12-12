package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateModuleVersionsRequest Request Object
type ListPrivateModuleVersionsRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有模块（private-module）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ModuleName string `json:"module_name"`

	// 私有模块（private-module）的唯一Id。  此Id由资源编排服务在生成模块的时候生成，为UUID。  由于私有模块名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有模块，删除，再重新创建一个同名私有模块。  对于团队并行开发，用户可能希望确保，当前我操作的私有模块就是我认为的那个，而不是其他队友删除后创建的同名私有模块。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有模块所对应的Id都不相同，更新不会影响Id。如果给予的module_id和当前模块的Id不一致，则返回400
	ModuleId *string `json:"module_id,omitempty"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListPrivateModuleVersionsRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListPrivateModuleVersionsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListPrivateModuleVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateModuleVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateModuleVersionsRequest", string(data)}, " ")
}

type ListPrivateModuleVersionsRequestSortKey struct {
	value string
}

type ListPrivateModuleVersionsRequestSortKeyEnum struct {
	CREATE_TIME ListPrivateModuleVersionsRequestSortKey
}

func GetListPrivateModuleVersionsRequestSortKeyEnum() ListPrivateModuleVersionsRequestSortKeyEnum {
	return ListPrivateModuleVersionsRequestSortKeyEnum{
		CREATE_TIME: ListPrivateModuleVersionsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListPrivateModuleVersionsRequestSortKey) Value() string {
	return c.value
}

func (c ListPrivateModuleVersionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateModuleVersionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListPrivateModuleVersionsRequestSortDir struct {
	value string
}

type ListPrivateModuleVersionsRequestSortDirEnum struct {
	ASC  ListPrivateModuleVersionsRequestSortDir
	DESC ListPrivateModuleVersionsRequestSortDir
}

func GetListPrivateModuleVersionsRequestSortDirEnum() ListPrivateModuleVersionsRequestSortDirEnum {
	return ListPrivateModuleVersionsRequestSortDirEnum{
		ASC: ListPrivateModuleVersionsRequestSortDir{
			value: "asc",
		},
		DESC: ListPrivateModuleVersionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPrivateModuleVersionsRequestSortDir) Value() string {
	return c.value
}

func (c ListPrivateModuleVersionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateModuleVersionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
