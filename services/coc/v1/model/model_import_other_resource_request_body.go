package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type ImportOtherResourceRequestBody struct {

	// **参数解释：** 上传的物理机/虚拟机/中间件设备下载模板excel（相关的设备信息）。 **约束限制：** 不涉及。 **取值范围：** 物理机/虚拟机/中间件设备下载模板excel（相关的设备信息）。 **默认取值：** 不涉及。
	File *def.FilePart `json:"file"`

	// **参数解释：** 导入类型。 **约束限制：** 不涉及。 **取值范围：** - vm：虚拟机。 - PM：物理机。 - Middleware：中间件设备。 **默认取值：** 不涉及。
	Type *def.MultiPart `json:"type"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度在[0,64]之间。 **默认取值：** 不涉及。
	RegionId *def.MultiPart `json:"region_id"`
}

func (o ImportOtherResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportOtherResourceRequestBody struct{}"
	}

	return strings.Join([]string{"ImportOtherResourceRequestBody", string(data)}, " ")
}

func (o *ImportOtherResourceRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
}

type ImportOtherResourceRequestBodyType struct {
	value string
}

type ImportOtherResourceRequestBodyTypeEnum struct {
	PM         ImportOtherResourceRequestBodyType
	VM         ImportOtherResourceRequestBodyType
	MIDDLEWARE ImportOtherResourceRequestBodyType
}

func GetImportOtherResourceRequestBodyTypeEnum() ImportOtherResourceRequestBodyTypeEnum {
	return ImportOtherResourceRequestBodyTypeEnum{
		PM: ImportOtherResourceRequestBodyType{
			value: "PM",
		},
		VM: ImportOtherResourceRequestBodyType{
			value: "VM",
		},
		MIDDLEWARE: ImportOtherResourceRequestBodyType{
			value: "Middleware",
		},
	}
}

func (c ImportOtherResourceRequestBodyType) Value() string {
	return c.value
}

func (c ImportOtherResourceRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportOtherResourceRequestBodyType) UnmarshalJSON(b []byte) error {
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
