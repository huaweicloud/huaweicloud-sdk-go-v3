package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveBrowsingHistoryResponse Response Object
type SaveBrowsingHistoryResponse struct {

	// uuid，仅当请求type为in时返回
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveBrowsingHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveBrowsingHistoryResponse struct{}"
	}

	return strings.Join([]string{"SaveBrowsingHistoryResponse", string(data)}, " ")
}
