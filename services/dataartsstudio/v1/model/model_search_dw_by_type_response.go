package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDwByTypeResponse Response Object
type SearchDwByTypeResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchDwByTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDwByTypeResponse struct{}"
	}

	return strings.Join([]string{"SearchDwByTypeResponse", string(data)}, " ")
}
