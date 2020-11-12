/*
 * CloudPipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type PipelineParameter struct {
	// 参数名称
	Name string `json:"name"`
	// 参数值
	Value string `json:"value"`
}

func (o PipelineParameter) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PipelineParameter", string(data)}, " ")
}
