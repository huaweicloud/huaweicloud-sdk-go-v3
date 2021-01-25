package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAppsV2Response struct {
	// 符合条件的APP总数
	Total *int32 `json:"total,omitempty"`
	// 本次查询返回的列表长度
	Size *int32 `json:"size,omitempty"`
	// APP列表
	Apps           *[]AppInfoWithBindNumResp `json:"apps,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAppsV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsV2Response struct{}"
	}

	return strings.Join([]string{"ListAppsV2Response", string(data)}, " ")
}
