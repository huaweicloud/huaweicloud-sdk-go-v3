package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type CreateUnscopeTokenByIdpInitiatedRequestBody struct {

	// 在IdP认证成功后返回的响应体。详情请参见：[获取联邦认证unscoped token(IdP initiated)](https://support.huaweicloud.com/api-iam/iam_02_0003.html)。
	SAMLResponse *def.MultiPart `json:"SAMLResponse"`
}

func (o CreateUnscopeTokenByIdpInitiatedRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnscopeTokenByIdpInitiatedRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUnscopeTokenByIdpInitiatedRequestBody", string(data)}, " ")
}

func (o *CreateUnscopeTokenByIdpInitiatedRequestBody) UnmarshalJSON(b []byte) error {
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
