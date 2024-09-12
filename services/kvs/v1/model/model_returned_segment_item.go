package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// ReturnedSegmentItem 采样段的起始终止主键对。
type ReturnedSegmentItem struct {

	// 采样段区间起始值。
	SegmentMinKey *bson.D `bson:"segment_min_key,omitempty"`

	// 采样段区间终止值。
	SegmentMaxKey *bson.D `bson:"segment_max_key,omitempty"`
}

func (o ReturnedSegmentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReturnedSegmentItem struct{}"
	}

	return strings.Join([]string{"ReturnedSegmentItem", string(data)}, " ")
}
