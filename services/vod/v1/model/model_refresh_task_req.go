package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RefreshTaskReq struct {

	// 完成态、删除态、发布态的媒资ID列表，一次性最多刷新10个媒资ID。
	AssetIdList *[]string `json:"asset_id_list,omitempty"`

	// 完成态、删除态、发布态的播放URL列表，一次性最多刷新10个URL。 单个URL的长度不能超过10240。
	Urls *[]string `json:"urls,omitempty"`
}

func (o RefreshTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshTaskReq struct{}"
	}

	return strings.Join([]string{"RefreshTaskReq", string(data)}, " ")
}
