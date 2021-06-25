package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListWatermarkTemplateRequest struct {
	// 水印模板配置id，一次最多10个<br/>

	Id *[]string `json:"id,omitempty"`
	// 分页编号。默认为0。指定id时该参数无效。<br/>

	Page *int32 `json:"page,omitempty"`
	// 每页记录数。默认为10，范围[1,100]。指定id时该参数无效。<br/>

	Size *int32 `json:"size,omitempty"`
}

func (o ListWatermarkTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateRequest", string(data)}, " ")
}
