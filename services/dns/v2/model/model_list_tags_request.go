package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// **参数解释：** 资源的类型。 **约束限制：** 不涉及。 **取值范围：** - DNS-public_zone：公网域名 - DNS-private_zone：内网域名 - DNS-public_recordset：公网记录集 - DNS-private_recordset：内网记录集 - DNS-ptr_record：反向解析  **默认取值：** 不涉及。
	ResourceType string `json:"resource_type"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}
