package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAssetCategoryResponse struct {
	// 分类返回值<br/>

	Body           *[]QueryCategoryRsp `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAssetCategoryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAssetCategoryResponse struct{}"
	}

	return strings.Join([]string{"ListAssetCategoryResponse", string(data)}, " ")
}
