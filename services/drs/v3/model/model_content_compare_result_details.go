package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ContentCompareResultDetails struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name"`

	// 该库的表的内容对比详情。
	ContentCompareDetail []ContentCompareDetail `json:"content_compare_detail"`

	// 内容对比结果详情总数。
	ContentCompareDetailCount int32 `json:"content_compare_detail_count"`

	// 该库的表的内容对比详情(无法对比的表)。
	ContentUncompareDetail *[]ContentCompareDetail `json:"content_uncompare_detail,omitempty"`

	// 内容对比结果详情总数(无法对比的表)。
	ContentUncompareDetailCount int32 `json:"content_uncompare_detail_count"`
}

func (o ContentCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultDetails struct{}"
	}

	return strings.Join([]string{"ContentCompareResultDetails", string(data)}, " ")
}
