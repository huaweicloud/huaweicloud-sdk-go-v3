package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessInfosResponse Response Object
type ListAccessInfosResponse struct {

	// accessinfo列表
	AccessInfos *[]AccessInfo `json:"access_infos,omitempty"`

	// accessinfo总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAccessInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAccessInfosResponse", string(data)}, " ")
}
