package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulAffectImagesResponse Response Object
type ListVulAffectImagesResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 受漏洞影响的镜像列表 **取值范围**: 不涉及
	DataList       *[]VulAffectImagesResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListVulAffectImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulAffectImagesResponse struct{}"
	}

	return strings.Join([]string{"ListVulAffectImagesResponse", string(data)}, " ")
}
