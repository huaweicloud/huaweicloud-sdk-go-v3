package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelRelationsResponse Response Object
type ListTableModelRelationsResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTableModelRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsResponse", string(data)}, " ")
}
