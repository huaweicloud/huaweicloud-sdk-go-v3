package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRefsListResponse Response Object
type ListRefsListResponse struct {

	// 分支/tag列表
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRefsListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRefsListResponse struct{}"
	}

	return strings.Join([]string{"ListRefsListResponse", string(data)}, " ")
}
