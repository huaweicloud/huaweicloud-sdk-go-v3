package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchFieldsForRelationResponse Response Object
type SearchFieldsForRelationResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchFieldsForRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFieldsForRelationResponse struct{}"
	}

	return strings.Join([]string{"SearchFieldsForRelationResponse", string(data)}, " ")
}
