package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSpaceAnalysisResponse struct {
	// 记录总数

	Total *int64 `json:"total,omitempty"`
	// 数据库对象列表

	DbObjects *[]DbObjectSpaceInfo `json:"db_objects,omitempty"`

	InstanceInfo   *InstanceSpaceInfo `json:"instance_info,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSpaceAnalysisResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSpaceAnalysisResponse struct{}"
	}

	return strings.Join([]string{"ListSpaceAnalysisResponse", string(data)}, " ")
}
