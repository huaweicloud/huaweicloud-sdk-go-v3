package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublisherResponse Response Object
type ShowPublisherResponse struct {

	// **参数解释**： 发布商详情。 **取值范围**： 不涉及。
	PublisherDetailMap map[string]PublisherVo `json:"publisher_detail_map,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o ShowPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublisherResponse struct{}"
	}

	return strings.Join([]string{"ShowPublisherResponse", string(data)}, " ")
}
