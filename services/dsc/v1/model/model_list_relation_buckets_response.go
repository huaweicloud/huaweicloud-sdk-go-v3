package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRelationBucketsResponse struct {

	// 关系信息总数
	Total *int32 `json:"total,omitempty"`

	// 关系信息列表
	BucketList     *[]RelationSimpleInfo `json:"bucket_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRelationBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListRelationBucketsResponse", string(data)}, " ")
}
