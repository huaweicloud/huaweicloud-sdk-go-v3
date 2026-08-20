package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopPathResponse Response Object
type ListCdnDomainTopPathResponse struct {

	// **参数解释：** 数据详情 **取值范围：** 不涉及
	TopPathSummary *[]TopPathSummary `json:"top_path_summary,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCdnDomainTopPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopPathResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopPathResponse", string(data)}, " ")
}
