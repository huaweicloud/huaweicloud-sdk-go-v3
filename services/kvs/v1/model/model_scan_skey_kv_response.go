package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// ScanSkeyKvResponse Response Object
type ScanSkeyKvResponse struct {

	// 返回的文档数量，为0不表示结束。 - 如果filtered_count超过500仍无匹配，则返回0。 -  长度：4
	ReturnedCount *int32 `bson:"returned_count,omitempty"`

	// 被过滤掉的文档数量。 - 长度：4
	FilteredCount *int32 `bson:"filtered_count,omitempty"`

	// 下次请求时的start_key。 > 该值为空时，表示指定范围或者指定filter条件所有kv已经返回。
	CursorSortKey *bson.D `bson:"cursor_sort_key,omitempty"`

	// 返回的kv列表，与scan_kv的kv_array相同。
	ReturnedKvItems *[]ReturnedKvItem `bson:"returned_kv_items,omitempty"`
	HttpStatusCode  int               `bson:"-"`
}

func (o ScanSkeyKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanSkeyKvResponse struct{}"
	}

	return strings.Join([]string{"ScanSkeyKvResponse", string(data)}, " ")
}
