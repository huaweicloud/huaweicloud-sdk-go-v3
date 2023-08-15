package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceRelationsDetailResponse Response Object
type ShowResourceRelationsDetailResponse struct {

	// 资源关系列表
	Relations *[]ResourceRelation `json:"relations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowResourceRelationsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRelationsDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceRelationsDetailResponse", string(data)}, " ")
}
