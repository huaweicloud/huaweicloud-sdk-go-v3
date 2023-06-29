package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResDatasourceRequestBody This is a auto create Body Object
type CreateResDatasourceRequestBody struct {

	// 数据源名称，1-64位字母、数字、下划线、中划线组合。
	DatasourceName string `json:"datasource_name"`

	SpecsConfig *SpecsConfig `json:"specs_config"`

	DataConfig *DataConfig `json:"data_config"`
}

func (o CreateResDatasourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResDatasourceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResDatasourceRequestBody", string(data)}, " ")
}
