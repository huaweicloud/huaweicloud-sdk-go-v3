package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAccessPointRequestBody 创建接入点请求体
type CreateAccessPointRequestBody struct {

	// **参数解释：** 接入点归属的可信密钥空间ID **约束限制：** UUID格式，满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 接入点的名称 **约束限制：** 满足正则表达式^[a-zA-Z0-9:/_-]{1,255}$ **取值范围：** 1-255 **默认取值：** 不涉及
	AccessPointName string `json:"access_point_name"`

	// **参数解释：** 接入点的类型 **约束限制：** 不涉及 **取值范围：** - 1：ECS - 2：CCE - 3：Custom **默认取值：** 不涉及
	Type CreateAccessPointRequestBodyType `json:"type"`

	// **参数解释：** 接入点的唯一标志 **约束限制：** ECS接入点填入ecs_id CCE接入点填入CCE集群公钥信息 Custom接入点无需填写，创建Custom接入点后，会生成一对密钥对，可以下载私钥，使用私钥签名，服务端验证签名 **取值范围：** 不涉及 **默认取值：** 不涉及
	Identity *string `json:"identity,omitempty"`

	// **参数解释：** 创建CCE接入点时必填，CCE集群ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释：** 接入点描述信息 **约束限制：** 不涉及 **取值范围：** 1-255 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`
}

func (o CreateAccessPointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessPointRequestBody", string(data)}, " ")
}

type CreateAccessPointRequestBodyType struct {
	value int32
}

type CreateAccessPointRequestBodyTypeEnum struct {
	E_1 CreateAccessPointRequestBodyType
	E_2 CreateAccessPointRequestBodyType
	E_3 CreateAccessPointRequestBodyType
}

func GetCreateAccessPointRequestBodyTypeEnum() CreateAccessPointRequestBodyTypeEnum {
	return CreateAccessPointRequestBodyTypeEnum{
		E_1: CreateAccessPointRequestBodyType{
			value: 1,
		}, E_2: CreateAccessPointRequestBodyType{
			value: 2,
		}, E_3: CreateAccessPointRequestBodyType{
			value: 3,
		},
	}
}

func (c CreateAccessPointRequestBodyType) Value() int32 {
	return c.value
}

func (c CreateAccessPointRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessPointRequestBodyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
