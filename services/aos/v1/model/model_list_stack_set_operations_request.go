package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListStackSetOperationsRequest Request Object
type ListStackSetOperationsRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 资源栈集的名称。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，再重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我认为的那个，而不是其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 过滤条件  * 与（AND）运算符使用逗号（，）定义 * 或（OR）运算符使用竖线（|）定义，OR运算符优先级高于AND运算符 * 不支持括号 * 过滤运算符仅支持双等号（==） * 过滤参数名及其值仅支持包含大小写英文、数字和下划线 * 过滤条件中禁止使用分号，如果有分号，则此条过滤会被忽略 * 一个过滤参数仅能与一个与条件相关，一个与条件中的多个或条件仅能与一个过滤参数相关
	Filter *string `json:"filter,omitempty"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListStackSetOperationsRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListStackSetOperationsRequestSortDir `json:"sort_dir,omitempty"`

	// 仅支持资源栈集权限模式为SERVICE_MANAGED时指定该参数。用于指定用户是以组织管理账号还是成员帐户中的服务委托管理员身份调用资源栈集。默认为SELF。 * 无论指定何种用户身份，创建或部署的资源栈集始终在组织管理账号名下。*   * `SELF` - 以组织管理账号身份调用。   * `DELEGATED_ADMIN` - 以服务委托管理员身份调用。用户的华为云账号必须在组织中已经被注册为”资源编排资源栈集服务“的委托管理员。
	CallIdentity *ListStackSetOperationsRequestCallIdentity `json:"call_identity,omitempty"`
}

func (o ListStackSetOperationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackSetOperationsRequest struct{}"
	}

	return strings.Join([]string{"ListStackSetOperationsRequest", string(data)}, " ")
}

type ListStackSetOperationsRequestSortKey struct {
	value string
}

type ListStackSetOperationsRequestSortKeyEnum struct {
	CREATE_TIME ListStackSetOperationsRequestSortKey
}

func GetListStackSetOperationsRequestSortKeyEnum() ListStackSetOperationsRequestSortKeyEnum {
	return ListStackSetOperationsRequestSortKeyEnum{
		CREATE_TIME: ListStackSetOperationsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListStackSetOperationsRequestSortKey) Value() string {
	return c.value
}

func (c ListStackSetOperationsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStackSetOperationsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListStackSetOperationsRequestSortDir struct {
	value string
}

type ListStackSetOperationsRequestSortDirEnum struct {
	ASC  ListStackSetOperationsRequestSortDir
	DESC ListStackSetOperationsRequestSortDir
}

func GetListStackSetOperationsRequestSortDirEnum() ListStackSetOperationsRequestSortDirEnum {
	return ListStackSetOperationsRequestSortDirEnum{
		ASC: ListStackSetOperationsRequestSortDir{
			value: "asc",
		},
		DESC: ListStackSetOperationsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListStackSetOperationsRequestSortDir) Value() string {
	return c.value
}

func (c ListStackSetOperationsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStackSetOperationsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListStackSetOperationsRequestCallIdentity struct {
	value string
}

type ListStackSetOperationsRequestCallIdentityEnum struct {
	SELF            ListStackSetOperationsRequestCallIdentity
	DELEGATED_ADMIN ListStackSetOperationsRequestCallIdentity
}

func GetListStackSetOperationsRequestCallIdentityEnum() ListStackSetOperationsRequestCallIdentityEnum {
	return ListStackSetOperationsRequestCallIdentityEnum{
		SELF: ListStackSetOperationsRequestCallIdentity{
			value: "SELF",
		},
		DELEGATED_ADMIN: ListStackSetOperationsRequestCallIdentity{
			value: "DELEGATED_ADMIN",
		},
	}
}

func (c ListStackSetOperationsRequestCallIdentity) Value() string {
	return c.value
}

func (c ListStackSetOperationsRequestCallIdentity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStackSetOperationsRequestCallIdentity) UnmarshalJSON(b []byte) error {
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
