package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentDiffDetailInfo 内容对比不一致数据的详情。
type ContentDiffDetailInfo struct {

	// 数量。
	Count *int64 `json:"count,omitempty"`

	// 对比不一致详情：只有源库存在。
	TargetMetaIsNull *int64 `json:"target_meta_is_null,omitempty"`

	// 对比不一致详情：只有目标库存在。
	SourceMetaIsNull *int64 `json:"source_meta_is_null,omitempty"`

	// 对比不一致详情：源和目标端均存在。
	SourceTargetMetaNotNull *int64 `json:"source_target_meta_not_null,omitempty"`

	// 信息列表。
	ContentsInfos *[]ContentDiffDetailVo `json:"contents_infos,omitempty"`
}

func (o ContentDiffDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDiffDetailInfo struct{}"
	}

	return strings.Join([]string{"ContentDiffDetailInfo", string(data)}, " ")
}
