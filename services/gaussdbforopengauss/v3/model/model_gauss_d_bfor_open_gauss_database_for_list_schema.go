package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussDatabaseForListSchema 数据库schema信息。
type GaussDBforOpenGaussDatabaseForListSchema struct {

	// schema名称。
	SchemaName string `json:"schema_name"`

	// schema所属用户。
	Owner string `json:"owner"`
}

func (o GaussDBforOpenGaussDatabaseForListSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussDatabaseForListSchema struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussDatabaseForListSchema", string(data)}, " ")
}
