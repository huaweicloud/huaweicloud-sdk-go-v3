package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunSearchResponse struct {

	// 搜索完成返回success。
	Result *string `json:"result,omitempty"`

	Data           *SearchRestInfo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RunSearchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchResponse struct{}"
	}

	return strings.Join([]string{"RunSearchResponse", string(data)}, " ")
}
