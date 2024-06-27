package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionModelPrimitiveTypeHolder struct {

	// 权限模型，定义了RFS操作资源栈集时所需委托的创建方式，枚举值，默认为SELF_MANAGED。用户可以使用创建资源栈集（CreateStackSet）API 指定该参数。该参数暂不支持更新。用户如果想要更新权限模型，可以通过先删除再创建同名资源栈集实现。   * `SELF_MANAGED` - 自我管理，基于部署需求，用户需要提前手动创建委托，既包含管理账号授权给RFS的委托，也包含成员账号授权给管理账号的委托。如果委托不存在或权限不足，创建资源栈集不会失败，创建资源栈实例时才会报错。   * `SERVICE_MANAGED` - 服务管理，基于Organization服务，RFS会自动创建部署Organization 成员账号时所需的全部 IAM 委托。用户需要提前在Organization可信服务列表中将”资源编排资源栈集服务“启用，且只有Organization的管理账号或”资源编排资源栈集服务“的委托管理员，才允许指定SERVICE_MANAGED创建资源栈集，否则会报错。
	PermissionModel *PermissionModelPrimitiveTypeHolderPermissionModel `json:"permission_model,omitempty"`
}

func (o PermissionModelPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionModelPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PermissionModelPrimitiveTypeHolder", string(data)}, " ")
}

type PermissionModelPrimitiveTypeHolderPermissionModel struct {
	value string
}

type PermissionModelPrimitiveTypeHolderPermissionModelEnum struct {
	SELF_MANAGED    PermissionModelPrimitiveTypeHolderPermissionModel
	SERVICE_MANAGED PermissionModelPrimitiveTypeHolderPermissionModel
}

func GetPermissionModelPrimitiveTypeHolderPermissionModelEnum() PermissionModelPrimitiveTypeHolderPermissionModelEnum {
	return PermissionModelPrimitiveTypeHolderPermissionModelEnum{
		SELF_MANAGED: PermissionModelPrimitiveTypeHolderPermissionModel{
			value: "SELF_MANAGED",
		},
		SERVICE_MANAGED: PermissionModelPrimitiveTypeHolderPermissionModel{
			value: "SERVICE_MANAGED",
		},
	}
}

func (c PermissionModelPrimitiveTypeHolderPermissionModel) Value() string {
	return c.value
}

func (c PermissionModelPrimitiveTypeHolderPermissionModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionModelPrimitiveTypeHolderPermissionModel) UnmarshalJSON(b []byte) error {
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
