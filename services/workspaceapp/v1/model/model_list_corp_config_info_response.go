package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCorpConfigInfoResponse Response Object
type ListCorpConfigInfoResponse struct {

	// 批量配置项列表。
	ConfigInfos    *[]CorpConfigInfo `json:"config_infos,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCorpConfigInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCorpConfigInfoResponse struct{}"
	}

	return strings.Join([]string{"ListCorpConfigInfoResponse", string(data)}, " ")
}
