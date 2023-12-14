package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenTsdbSchema 转储OpenTSDB时必选，与“cloudtable_schema”二选一，表示CloudTable集群OpenTSDB数据的Schema配置。用于将通道内的JSON数据进行格式转换并导入Cloudtable的OpenTSDB。
type OpenTsdbSchema struct {

	// CloudTable集群OpenTSDB数据metric的Schema配置，用于将通道内的JSON数据进行格式转换生成OpenTSDB数据的metric。
	Metric []OpenTsdbMetric `json:"metric"`

	Timestamp *OpenTsdbTimestamp `json:"timestamp"`

	Value *OpenTsdbValue `json:"value"`

	// CloudTable集群OpenTSDB数据tags的Schema配置，用于将通道内的JSON数据进行格式转换生成OpenTSDB数据的tags。
	Tags []OpenTsdbTags `json:"tags"`
}

func (o OpenTsdbSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTsdbSchema struct{}"
	}

	return strings.Join([]string{"OpenTsdbSchema", string(data)}, " ")
}
