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

// RegisterCluster 集群信息，包括期望状态（例如同步模式），和集群状态（例如版本、Node状态信息及资源统计信息等）。
type RegisterCluster struct {

	// 资源类型。注册集群必须填写为Cluster。
	Kind *def.MultiPart `json:"kind"`

	// API版本信息。现版本仅为v1。
	ApiVersion *def.MultiPart `json:"apiVersion"`

	Metadata *RegisterClusterMetadata `json:"metadata"`

	Spec *RegisterClusterSpec `json:"spec"`
}

func (o RegisterCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterCluster struct{}"
	}

	return strings.Join([]string{"RegisterCluster", string(data)}, " ")
}

func (o *RegisterCluster) UnmarshalJSON(b []byte) error {
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
