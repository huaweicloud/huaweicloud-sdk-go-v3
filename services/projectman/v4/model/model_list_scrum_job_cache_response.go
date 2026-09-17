package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScrumJobCacheResponse Response Object
type ListScrumJobCacheResponse struct {
	Result *ListCacheDatasResposeResult `json:"result,omitempty"`

	// **参数解释：** 查询缓存的返回状态。 **取值范围：** success：返回成功。 error：返回失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListScrumJobCacheResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScrumJobCacheResponse struct{}"
	}

	return strings.Join([]string{"ListScrumJobCacheResponse", string(data)}, " ")
}
