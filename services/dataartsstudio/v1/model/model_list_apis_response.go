package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisResponse Response Object
type ListApisResponse struct {

	// API总数量
	Total *int32 `json:"total,omitempty"`

	// API列表
	Records        *[]ApiOverview `json:"records,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListApisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisResponse struct{}"
	}

	return strings.Join([]string{"ListApisResponse", string(data)}, " ")
}
