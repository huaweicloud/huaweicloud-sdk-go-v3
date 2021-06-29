package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAssetListResponse struct {
	// 媒资总数  **说明：** 暂只能统计2万个媒资，若您需要查询具体的媒资总数，请提交工单申请。

	Total *int32 `json:"total,omitempty"`
	// 媒资列表

	Assets         *[]AssetSummary `json:"assets,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAssetListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAssetListResponse struct{}"
	}

	return strings.Join([]string{"ListAssetListResponse", string(data)}, " ")
}
