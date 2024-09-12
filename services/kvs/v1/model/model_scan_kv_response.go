package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// ScanKvResponse Response Object
type ScanKvResponse struct {

	// 返回的文档数量，为0不表示结束。 > 如果filtered_count超过500仍无匹配，则返回0。
	ReturnedCount *int32 `bson:"returned_count,omitempty"`

	// 被过滤掉的文档数量。
	FilteredCount *int32 `bson:"filtered_count,omitempty"`

	// 下次请求时的start_key，该值为空时，表示指定范围或者指定filter条件所有kv已经返回。
	CursorKey *bson.D `bson:"cursor_key,omitempty"`

	// 返回的kv数据。
	ReturnedKvItems *[]ReturnedKvItem `bson:"returned_kv_items,omitempty"`

	// 采样段区间列表。
	ReturnedSegmentItems *[]ReturnedSegmentItem `bson:"returned_segment_items,omitempty"`
	HttpStatusCode       int                    `bson:"-"`
}

func (o ScanKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanKvResponse struct{}"
	}

	return strings.Join([]string{"ScanKvResponse", string(data)}, " ")
}
