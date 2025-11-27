package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationViewResponse Response Object
type ListApplicationViewResponse struct {

	// 查询应用、组件、分组名称列表。
	Data           *[]ApplicationViewQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListApplicationViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationViewResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationViewResponse", string(data)}, " ")
}
