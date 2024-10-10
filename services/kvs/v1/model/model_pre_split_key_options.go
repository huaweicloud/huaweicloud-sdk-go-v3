package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// PreSplitKeyOptions 按照设定的键值前缀进行预分裂。
type PreSplitKeyOptions struct {

	// 在range分区模式有效，最大10个。
	RangeSplitPoints *[]bson.D `bson:"range_split_points,omitempty"`
}

func (o PreSplitKeyOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreSplitKeyOptions struct{}"
	}

	return strings.Join([]string{"PreSplitKeyOptions", string(data)}, " ")
}
