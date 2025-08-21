package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataMapLineageRequestBody 导入血缘请求体。
type ImportDataMapLineageRequestBody struct {

	// 上游血缘sourceId。
	InputSourceId string `json:"input_source_id"`

	// 下游血缘sourceId。
	OutputSourceId string `json:"output_source_id"`

	// 血缘信息。
	LineageConfigs []SubNodeLineageConfig `json:"lineage_configs"`
}

func (o ImportDataMapLineageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataMapLineageRequestBody struct{}"
	}

	return strings.Join([]string{"ImportDataMapLineageRequestBody", string(data)}, " ")
}
