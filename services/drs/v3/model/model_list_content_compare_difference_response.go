package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContentCompareDifferenceResponse Response Object
type ListContentCompareDifferenceResponse struct {

	// 总数量。
	Count *int64 `json:"count,omitempty"`

	// 对比不一致详情数量：只有源库存在。
	TargetMetaIsNull *int64 `json:"target_meta_is_null,omitempty"`

	// 对比不一致详情数量：只有目标库存在。
	SourceMetaIsNull *int64 `json:"source_meta_is_null,omitempty"`

	// 对比不一致详情数量：源和目标端均存在。
	SourceTargetMetaNotNull *int64 `json:"source_target_meta_not_null,omitempty"`

	// 详细内容信息列表。
	ContentsInfos  *[]CompareJobContentDetailInfo `json:"contents_infos,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListContentCompareDifferenceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContentCompareDifferenceResponse struct{}"
	}

	return strings.Join([]string{"ListContentCompareDifferenceResponse", string(data)}, " ")
}
