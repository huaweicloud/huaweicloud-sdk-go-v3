package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCaTagsRequest Request Object
type BatchDeleteCaTagsRequest struct {

	// 所需要批量删除标签的CA证书ID。
	CaId string `json:"ca_id"`

	Body *BatchOperateTagRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteCaTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCaTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCaTagsRequest", string(data)}, " ")
}
