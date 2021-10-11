package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApiVersionsV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的API历史版本列表

	ApiVersions    *[]ApiVersionResp `json:"api_versions,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListApiVersionsV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionsV2Response struct{}"
	}

	return strings.Join([]string{"ListApiVersionsV2Response", string(data)}, " ")
}
