package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteImageWhiteListsRequestBody struct {

	// 白名单ID列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o DeleteImageWhiteListsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageWhiteListsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteImageWhiteListsRequestBody", string(data)}, " ")
}
