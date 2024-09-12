package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// PreSplitKeyOptions 在hash分区时，预分裂分区数量。
type PreSplitKeyOptions struct {

	// 在hash分区时，预分裂分区数量。
	HashCount *int32 `bson:"hash_count,omitempty"`

	// 在range分区模式有效，最大1000个，与\"hash_count\"二选一。
	RangeSplitPoints *[]bson.D `bson:"range_split_points,omitempty"`
}

func (o PreSplitKeyOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreSplitKeyOptions struct{}"
	}

	return strings.Join([]string{"PreSplitKeyOptions", string(data)}, " ")
}
