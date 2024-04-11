package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryRelationshipResponse Response Object
type ListQueryRelationshipResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]RelationModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryRelationshipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryRelationshipResponse struct{}"
	}

	return strings.Join([]string{"ListQueryRelationshipResponse", string(data)}, " ")
}
