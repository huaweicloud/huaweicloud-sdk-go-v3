package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventAttCkResponse Response Object
type ListEventAttCkResponse struct {

	// 各种攻击阶段详情
	DataList       *[]EventAttCkDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListEventAttCkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventAttCkResponse struct{}"
	}

	return strings.Join([]string{"ListEventAttCkResponse", string(data)}, " ")
}
