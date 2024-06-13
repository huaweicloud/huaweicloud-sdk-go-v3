package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteAssetActionResponse Response Object
type BatchExecuteAssetActionResponse struct {

	// 批量操作结果
	Results *[]AssetActionResult `json:"results,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchExecuteAssetActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteAssetActionResponse struct{}"
	}

	return strings.Join([]string{"BatchExecuteAssetActionResponse", string(data)}, " ")
}
