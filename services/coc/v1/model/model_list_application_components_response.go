package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationComponentsResponse Response Object
type ListApplicationComponentsResponse struct {

	// 组件查询信息列表。
	Data           *[]ComponentQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListApplicationComponentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationComponentsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationComponentsResponse", string(data)}, " ")
}
