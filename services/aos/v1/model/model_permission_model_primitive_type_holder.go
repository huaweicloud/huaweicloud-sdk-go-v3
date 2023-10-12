package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionModelPrimitiveTypeHolder struct {

	// 权限模型，定义了RFS操作资源栈集时所需委托的创建方式，枚举值     * `SERVICE_MANAGED` - 用户在自己的Organization中将RFS设置为可信服务后，才可以使用此种模式创建资源栈集。用户不需要手动创建委托，基于Organization，RFS会自动帮用户创建委托。只有Organization组织管理员或委托管理员才可以在SERVICE_MANAGED下创建资源栈集。     * `SELF_MANAGED` - 基于部署需求，用户需要提前手动创建委托，既包含管理账号给RFS的委托，也包含成员账号创建给管理账号的委托。如果委托不存在或错误，创建资源栈集不会失败，部署资源栈集或部署资源栈实例的时候才会报错。
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
	SELF_MANAGED PermissionModelPrimitiveTypeHolderPermissionModel
}

func GetPermissionModelPrimitiveTypeHolderPermissionModelEnum() PermissionModelPrimitiveTypeHolderPermissionModelEnum {
	return PermissionModelPrimitiveTypeHolderPermissionModelEnum{
		SELF_MANAGED: PermissionModelPrimitiveTypeHolderPermissionModel{
			value: "SELF_MANAGED",
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
