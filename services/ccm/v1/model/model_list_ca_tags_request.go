package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaTagsRequest Request Object
type ListCaTagsRequest struct {

	// 所需要查询标签列表的CA证书ID。
	CaId string `json:"ca_id"`
}

func (o ListCaTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCaTagsRequest", string(data)}, " ")
}
