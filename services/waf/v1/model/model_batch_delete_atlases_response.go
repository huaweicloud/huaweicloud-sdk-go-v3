package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAtlasesResponse Response Object
type BatchDeleteAtlasesResponse struct {

	// 删除成功的图库ID列表
	Deleted *[]string `json:"deleted,omitempty"`

	// 删除失败的图库列表
	Failed         *[]FailedAtlas `json:"failed,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchDeleteAtlasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAtlasesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAtlasesResponse", string(data)}, " ")
}
