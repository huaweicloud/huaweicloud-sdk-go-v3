package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRegistriesResponse Response Object
type ListInstanceRegistriesResponse struct {

	// 同步仓库列表
	Registries *[]Registry `json:"registries,omitempty"`

	// 同步仓库总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceRegistriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRegistriesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRegistriesResponse", string(data)}, " ")
}
