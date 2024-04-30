package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStackSetMetadataResponse Response Object
type ShowStackSetMetadataResponse struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 资源栈集的描述。可用于客户识别自己的资源栈集。
	StackSetDescription *string `json:"stack_set_description,omitempty"`

	// 初始化资源栈描述。可用于客户识别被资源栈集所管理的资源栈。  资源栈集下的资源栈仅在创建时统一使用该描述。客户想要更新初始化资源栈描述，可以通过UpdateStackSet API。  后续更新资源栈集描述将不会同步更新已管理的资源栈描述。
	InitialStackDescription *string `json:"initial_stack_description,omitempty"`

	// 权限模型，定义了RFS操作资源栈集时所需委托的创建方式，枚举值    * `SELF_MANAGED` - 基于部署需求，用户需要提前手动创建委托，既包含管理账号给RFS的委托，也包含成员账号创建给管理账号的委托。如果委托不存在或错误，创建资源栈集不会失败，部署资源栈集或部署资源栈实例的时候才会报错。
	PermissionModel *ShowStackSetMetadataResponsePermissionModel `json:"permission_model,omitempty"`

	// 管理委托名称  资源编排服务使用该委托获取成员账号委托给管理账号的权限。该委托中必须含有iam:tokens:assume权限，用以后续获取被管理委托凭证。如果不包含，则会在新增或者部署实例时报错。  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予administration_agency_urn，administration_agency_name只支持接收委托名称，如果给予了信任委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。  [[创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)](tag:hws) [[创建委托及授权方式](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)](tag:hws_hk) [[创建委托及授权方式](https://support.huaweicloud.com/eu/usermanual-iam/iam_06_0002.html)](tag:hws_eu)
	AdministrationAgencyName *string `json:"administration_agency_name,omitempty"`

	// 被管理的委托名称。  资源编排服务会使用该委托获取实际部署资源所需要的权限  不同成员账号委托给管理账号的委托名称需要保持一致。暂不支持根据不同provider定义不同委托权限  当用户定义SELF_MANAGED权限类型时，必须指定该参数。当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400  [[创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)](tag:hws) [[创建委托及授权方式](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)](tag:hws_hk) [[创建委托及授权方式](https://support.huaweicloud.com/eu/usermanual-iam/iam_06_0002.html)](tag:hws_eu)
	ManagedAgencyName *string `json:"managed_agency_name,omitempty"`

	// 资源栈集的状态     * `IDLE` - 资源栈集空闲 * `OPERATION_IN_PROGRESS` - 资源栈集操作中 * `DEACTIVATED` - 资源栈集禁用
	Status *ShowStackSetMetadataResponseStatus `json:"status,omitempty"`

	// vars_uri对应的文件内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`

	// HCL参数文件的内容。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  * vars_body使用HCL的tfvars格式，用户可以将“.tfvars”中的内容提交到vars_body中  * 资源编排服务支持vars_body和vars_uri，如果以上两种方式中声名了同一个变量，将报错400  * 如果vars_body过大，可以使用vars_uri  * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储对应的vars_body。
	VarsBody *string `json:"vars_body,omitempty"`

	// 资源栈集的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈集的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`

	// 管理委托URN  资源编排服务使用该委托获取成员账号委托给管理账号的权限。该委托中必须含有sts:tokens:assume权限，用以后续获取被管理委托凭证。如果不包含，则会在新增或者部署实例时报错。  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予administration_agency_urn，administration_agency_name只支持接收委托名称，如果给予了信任委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。
	AdministrationAgencyUrn *string `json:"administration_agency_urn,omitempty"`

	ManagedOperation *ManagedOperation `json:"managed_operation,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowStackSetMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStackSetMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowStackSetMetadataResponse", string(data)}, " ")
}

type ShowStackSetMetadataResponsePermissionModel struct {
	value string
}

type ShowStackSetMetadataResponsePermissionModelEnum struct {
	SELF_MANAGED ShowStackSetMetadataResponsePermissionModel
}

func GetShowStackSetMetadataResponsePermissionModelEnum() ShowStackSetMetadataResponsePermissionModelEnum {
	return ShowStackSetMetadataResponsePermissionModelEnum{
		SELF_MANAGED: ShowStackSetMetadataResponsePermissionModel{
			value: "SELF_MANAGED",
		},
	}
}

func (c ShowStackSetMetadataResponsePermissionModel) Value() string {
	return c.value
}

func (c ShowStackSetMetadataResponsePermissionModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStackSetMetadataResponsePermissionModel) UnmarshalJSON(b []byte) error {
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

type ShowStackSetMetadataResponseStatus struct {
	value string
}

type ShowStackSetMetadataResponseStatusEnum struct {
	IDLE                  ShowStackSetMetadataResponseStatus
	OPERATION_IN_PROGRESS ShowStackSetMetadataResponseStatus
	DEACTIVATED           ShowStackSetMetadataResponseStatus
}

func GetShowStackSetMetadataResponseStatusEnum() ShowStackSetMetadataResponseStatusEnum {
	return ShowStackSetMetadataResponseStatusEnum{
		IDLE: ShowStackSetMetadataResponseStatus{
			value: "IDLE",
		},
		OPERATION_IN_PROGRESS: ShowStackSetMetadataResponseStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		DEACTIVATED: ShowStackSetMetadataResponseStatus{
			value: "DEACTIVATED",
		},
	}
}

func (c ShowStackSetMetadataResponseStatus) Value() string {
	return c.value
}

func (c ShowStackSetMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStackSetMetadataResponseStatus) UnmarshalJSON(b []byte) error {
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
