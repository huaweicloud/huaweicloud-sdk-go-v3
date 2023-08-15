package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCatalogsResponse Response Object
type SearchCatalogsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchCatalogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCatalogsResponse struct{}"
	}

	return strings.Join([]string{"SearchCatalogsResponse", string(data)}, " ")
}
