package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDependenciesResponse Response Object
type ListDependenciesResponse struct {

	// 依赖包列表
	Dependencies *[]ListDependenciesResult `json:"dependencies,omitempty"`

	// 下次读取位置
	NextMarker *int64 `json:"next_marker,omitempty"`

	// 依赖包总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDependenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDependenciesResponse struct{}"
	}

	return strings.Join([]string{"ListDependenciesResponse", string(data)}, " ")
}
