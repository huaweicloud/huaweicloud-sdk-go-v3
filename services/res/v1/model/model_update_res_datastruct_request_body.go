package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateResDatastructRequestBody struct {

	// 数据源名称:，1-64位字母、数字、下划线、中划线组合。
	Name string `json:"name" xml:"name"`

	DataConfig *DataConfig `json:"data_config" xml:"data_config"`

	SpecsConfig *SpecsConfig `json:"specs_config" xml:"specs_config"`
}

func (o UpdateResDatastructRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatastructRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResDatastructRequestBody", string(data)}, " ")
}
