package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCriteriasResponse struct {

	// 快速查询id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCriteriasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCriteriasResponse struct{}"
	}

	return strings.Join([]string{"ListCriteriasResponse", string(data)}, " ")
}
