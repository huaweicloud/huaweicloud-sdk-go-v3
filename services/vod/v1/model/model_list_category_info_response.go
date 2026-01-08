package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoryInfoResponse Response Object
type ListCategoryInfoResponse struct {

	// 分类返回值
	Body           *[]QueryCategoryInfoRsp `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListCategoryInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoryInfoResponse struct{}"
	}

	return strings.Join([]string{"ListCategoryInfoResponse", string(data)}, " ")
}
