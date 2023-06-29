package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchCriteriasResponse Response Object
type CreateSearchCriteriasResponse struct {

	// 快速查询id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSearchCriteriasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchCriteriasResponse struct{}"
	}

	return strings.Join([]string{"CreateSearchCriteriasResponse", string(data)}, " ")
}
