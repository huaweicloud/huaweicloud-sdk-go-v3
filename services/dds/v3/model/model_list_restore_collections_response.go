package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreCollectionsResponse Response Object
type ListRestoreCollectionsResponse struct {

	// 集合总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 集合列表，列表中每个元素表示一个集合。
	Collections    *[]string `json:"collections,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreCollectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreCollectionsResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreCollectionsResponse", string(data)}, " ")
}
