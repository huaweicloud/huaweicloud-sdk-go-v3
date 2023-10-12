package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateStackSetRequestBody struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给与的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集的描述。可用于客户识别自己的资源栈集。
	StackSetDescription *string `json:"stack_set_description,omitempty"`

	// 初始化资源栈描述。可用于客户识别被资源栈集所管理的资源栈。  资源栈集下的资源栈仅在创建时统一使用该描述。客户想要更新初始化资源栈描述，可以通过UpdateStackSet API。  后续更新资源栈集描述将不会同步更新已管理的资源栈描述。
	InitialStackDescription *string `json:"initial_stack_description,omitempty"`

	// 权限模型，定义了RFS操作资源栈集时所需委托的创建方式，枚举值     * `SERVICE_MANAGED` - 用户在自己的Organization中将RFS设置为可信服务后，才可以使用此种模式创建资源栈集。用户不需要手动创建委托，基于Organization，RFS会自动帮用户创建委托。只有Organization组织管理员或委托管理员才可以在SERVICE_MANAGED下创建资源栈集。     * `SELF_MANAGED` - 基于部署需求，用户需要提前手动创建委托，既包含管理账号给RFS的委托，也包含成员账号创建给管理账号的委托。如果委托不存在或错误，创建资源栈集不会失败，部署资源栈集或部署资源栈实例的时候才会报错。
	PermissionModel *UpdateStackSetRequestBodyPermissionModel `json:"permission_model,omitempty"`

	// 管理委托名称  资源编排服务使用该委托获取成员账号委托给管理账号的权限  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用v5委托时给与administration_agency_urn，administration_agency_name只支持接收v3委托名称，若给与了v5委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。
	AdministrationAgencyName *string `json:"administration_agency_name,omitempty"`

	// 被管理的委托名称。  资源编排服务会使用该委托获取实际部署资源所需要的权限  不同成员账号委托给管理账号的委托名称需要保持一致。暂不支持根据不同provider定义不同委托权限  当用户定义SELF_MANAGED权限类型时，必须指定该参数。当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400
	ManagedAgencyName *string `json:"managed_agency_name,omitempty"`

	// 管理委托URN  资源编排服务使用该委托获取成员账号委托给管理账号的权限  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用v5委托时给与administration_agency_urn，administration_agency_name只支持接收普通委托名称，若给与了v5委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。
	AdministrationAgencyUrn *string `json:"administration_agency_urn,omitempty"`
}

func (o UpdateStackSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStackSetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateStackSetRequestBody", string(data)}, " ")
}

type UpdateStackSetRequestBodyPermissionModel struct {
	value string
}

type UpdateStackSetRequestBodyPermissionModelEnum struct {
	SELF_MANAGED UpdateStackSetRequestBodyPermissionModel
}

func GetUpdateStackSetRequestBodyPermissionModelEnum() UpdateStackSetRequestBodyPermissionModelEnum {
	return UpdateStackSetRequestBodyPermissionModelEnum{
		SELF_MANAGED: UpdateStackSetRequestBodyPermissionModel{
			value: "SELF_MANAGED",
		},
	}
}

func (c UpdateStackSetRequestBodyPermissionModel) Value() string {
	return c.value
}

func (c UpdateStackSetRequestBodyPermissionModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateStackSetRequestBodyPermissionModel) UnmarshalJSON(b []byte) error {
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
