package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DbObject 数据库对象信息体。实时迁移、实时同步需要迁移或同步的库或者表，支持实时同步场景对数据库对象进行加工，即可以为数据库对象添加过滤规则、高级设置、列加工、附加列等。 数据加工相关具体约束参考：https://support.huaweicloud.com/realtimesyn-drs/drs_03_0035.html
type DbObject struct {

	// 数据库对象迁移或同步范围。取值： - all：全部迁移。 - database：库级迁移或同步。 - table：表级迁移或同步。
	ObjectScope DbObjectObjectScope `json:"object_scope"`

	TargetRootDb *TargetRootDb `json:"target_root_db,omitempty"`

	// 数据库对象迁移或同步信息，object_scope为all时不填，为库级或表级时必填。
	ObjectInfo map[string]DatabaseObject `json:"object_info,omitempty"`
}

func (o DbObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObject struct{}"
	}

	return strings.Join([]string{"DbObject", string(data)}, " ")
}

type DbObjectObjectScope struct {
	value string
}

type DbObjectObjectScopeEnum struct {
	ALL      DbObjectObjectScope
	DATABASE DbObjectObjectScope
	TABLE    DbObjectObjectScope
}

func GetDbObjectObjectScopeEnum() DbObjectObjectScopeEnum {
	return DbObjectObjectScopeEnum{
		ALL: DbObjectObjectScope{
			value: "all",
		},
		DATABASE: DbObjectObjectScope{
			value: "database",
		},
		TABLE: DbObjectObjectScope{
			value: "table",
		},
	}
}

func (c DbObjectObjectScope) Value() string {
	return c.value
}

func (c DbObjectObjectScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbObjectObjectScope) UnmarshalJSON(b []byte) error {
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
