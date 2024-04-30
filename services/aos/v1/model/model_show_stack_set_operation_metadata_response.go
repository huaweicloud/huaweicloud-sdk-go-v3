package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStackSetOperationMetadataResponse Response Object
type ShowStackSetOperationMetadataResponse struct {

	// 资源栈集操作（stack_set_operation）的唯一Id。  此ID由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	StackSetOperationId *string `json:"stack_set_operation_id,omitempty"`

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 资源栈集操作状态   * `QUEUE_IN_PROGRESS` - 正在排队   * `OPERATION_IN_PROGRESS` - 正在操作   * `OPERATION_COMPLETE` - 操作完成   * `OPERATION_FAILED` - 操作失败   * `STOP_IN_PROGRESS` - 正在停止   * `STOP_COMPLETE` - 停止完成   * `STOP_FAILED` - 停止失败
	Status *ShowStackSetOperationMetadataResponseStatus `json:"status,omitempty"`

	// 资源栈集操作失败时会展示此次操作失败的原因，例如，资源栈实例部署或删除失败个数超过上限或资源栈集操作超时。  如果需要查看详细失败信息，可通过ListStackInstances API获取查看资源栈实例的status_message。
	StatusMessage *string `json:"status_message,omitempty"`

	// 用户当前的操作   * `CREATE_STACK_INSTANCES` - 创建资源栈实例   * `DELETE_STACK_INSTANCES` - 删除资源栈实例   * `DEPLOY_STACK_SET` - 部署资源栈集   * `DEPLOY_STACK_INSTANCES` - 部署资源栈实例   * `UPDATE_STACK_INSTANCES` - 更新资源栈实例
	Action *ShowStackSetOperationMetadataResponseAction `json:"action,omitempty"`

	// 管理委托名称  资源编排服务使用该委托获取成员账号委托给管理账号的权限。该委托中必须含有iam:tokens:assume权限，用以后续获取被管理委托凭证。如果不包含，则会在新增或者部署实例时报错。  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予administration_agency_urn，administration_agency_name只支持接收委托名称，如果给予了信任委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。  [[创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)](tag:hws) [[创建委托及授权方式](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)](tag:hws_hk) [[创建委托及授权方式](https://support.huaweicloud.com/eu/usermanual-iam/iam_06_0002.html)](tag:hws_eu)
	AdministrationAgencyName *string `json:"administration_agency_name,omitempty"`

	// 管理委托URN  资源编排服务使用该委托获取成员账号委托给管理账号的权限。该委托中必须含有sts:tokens:assume权限，用以后续获取被管理委托凭证。如果不包含，则会在新增或者部署实例时报错。  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予administration_agency_urn，administration_agency_name只支持接收委托名称，如果给予了信任委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。
	AdministrationAgencyUrn *string `json:"administration_agency_urn,omitempty"`

	// 被管理的委托名称。  资源编排服务会使用该委托获取实际部署资源所需要的权限  不同成员账号委托给管理账号的委托名称需要保持一致。暂不支持根据不同provider定义不同委托权限  当用户定义SELF_MANAGED权限类型时，必须指定该参数。当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400  [[创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)](tag:hws) [[创建委托及授权方式](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)](tag:hws_hk) [[创建委托及授权方式](https://support.huaweicloud.com/eu/usermanual-iam/iam_06_0002.html)](tag:hws_eu)
	ManagedAgencyName *string `json:"managed_agency_name,omitempty"`

	DeploymentTargets *DeploymentTargets `json:"deployment_targets"`

	// 资源栈集操作的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈集操作的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`

	OperationPreferences *OperationPreferences `json:"operation_preferences,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowStackSetOperationMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStackSetOperationMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowStackSetOperationMetadataResponse", string(data)}, " ")
}

type ShowStackSetOperationMetadataResponseStatus struct {
	value string
}

type ShowStackSetOperationMetadataResponseStatusEnum struct {
	QUEUE_IN_PROGRESS     ShowStackSetOperationMetadataResponseStatus
	OPERATION_IN_PROGRESS ShowStackSetOperationMetadataResponseStatus
	OPERATION_COMPLETE    ShowStackSetOperationMetadataResponseStatus
	OPERATION_FAILED      ShowStackSetOperationMetadataResponseStatus
	STOP_IN_PROGRESS      ShowStackSetOperationMetadataResponseStatus
	STOP_COMPLETE         ShowStackSetOperationMetadataResponseStatus
	STOP_FAILED           ShowStackSetOperationMetadataResponseStatus
}

func GetShowStackSetOperationMetadataResponseStatusEnum() ShowStackSetOperationMetadataResponseStatusEnum {
	return ShowStackSetOperationMetadataResponseStatusEnum{
		QUEUE_IN_PROGRESS: ShowStackSetOperationMetadataResponseStatus{
			value: "QUEUE_IN_PROGRESS",
		},
		OPERATION_IN_PROGRESS: ShowStackSetOperationMetadataResponseStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		OPERATION_COMPLETE: ShowStackSetOperationMetadataResponseStatus{
			value: "OPERATION_COMPLETE",
		},
		OPERATION_FAILED: ShowStackSetOperationMetadataResponseStatus{
			value: "OPERATION_FAILED",
		},
		STOP_IN_PROGRESS: ShowStackSetOperationMetadataResponseStatus{
			value: "STOP_IN_PROGRESS",
		},
		STOP_COMPLETE: ShowStackSetOperationMetadataResponseStatus{
			value: "STOP_COMPLETE",
		},
		STOP_FAILED: ShowStackSetOperationMetadataResponseStatus{
			value: "STOP_FAILED",
		},
	}
}

func (c ShowStackSetOperationMetadataResponseStatus) Value() string {
	return c.value
}

func (c ShowStackSetOperationMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStackSetOperationMetadataResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowStackSetOperationMetadataResponseAction struct {
	value string
}

type ShowStackSetOperationMetadataResponseActionEnum struct {
	CREATE_STACK_INSTANCES ShowStackSetOperationMetadataResponseAction
	DELETE_STACK_INSTANCES ShowStackSetOperationMetadataResponseAction
	DEPLOY_STACK_SET       ShowStackSetOperationMetadataResponseAction
	DEPLOY_STACK_INSTANCES ShowStackSetOperationMetadataResponseAction
	UPDATE_STACK_INSTANCES ShowStackSetOperationMetadataResponseAction
}

func GetShowStackSetOperationMetadataResponseActionEnum() ShowStackSetOperationMetadataResponseActionEnum {
	return ShowStackSetOperationMetadataResponseActionEnum{
		CREATE_STACK_INSTANCES: ShowStackSetOperationMetadataResponseAction{
			value: "CREATE_STACK_INSTANCES",
		},
		DELETE_STACK_INSTANCES: ShowStackSetOperationMetadataResponseAction{
			value: "DELETE_STACK_INSTANCES",
		},
		DEPLOY_STACK_SET: ShowStackSetOperationMetadataResponseAction{
			value: "DEPLOY_STACK_SET",
		},
		DEPLOY_STACK_INSTANCES: ShowStackSetOperationMetadataResponseAction{
			value: "DEPLOY_STACK_INSTANCES",
		},
		UPDATE_STACK_INSTANCES: ShowStackSetOperationMetadataResponseAction{
			value: "UPDATE_STACK_INSTANCES",
		},
	}
}

func (c ShowStackSetOperationMetadataResponseAction) Value() string {
	return c.value
}

func (c ShowStackSetOperationMetadataResponseAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStackSetOperationMetadataResponseAction) UnmarshalJSON(b []byte) error {
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
