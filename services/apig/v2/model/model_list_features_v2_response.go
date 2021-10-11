package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFeaturesV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 实例特性列表

	Features       *[]FeatureInfo `json:"features,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFeaturesV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFeaturesV2Response struct{}"
	}

	return strings.Join([]string{"ListFeaturesV2Response", string(data)}, " ")
}
