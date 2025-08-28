package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageHandleAffectVulnerabilitiesResponse Response Object
type ListImageHandleAffectVulnerabilitiesResponse struct {

	// **参数解释**: 镜像-漏洞总条数 **取值范围**: 最小值0，最大值2147483547
	TotalVulNum *int32 `json:"total_vul_num,omitempty"`

	// **参数解释**: 漏洞个数 **取值范围**: 最小值0，最大值2147483547
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 镜像个数 **取值范围**: 最小值0，最大值2147483547
	ImageNum *int32 `json:"image_num,omitempty"`

	// **参数解释**: 镜像漏洞列表 **取值范围**: 最小值0，最大值5000
	DataList       *[]ImageVulResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListImageHandleAffectVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageHandleAffectVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListImageHandleAffectVulnerabilitiesResponse", string(data)}, " ")
}
