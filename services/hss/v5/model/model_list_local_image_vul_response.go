package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLocalImageVulResponse Response Object
type ListLocalImageVulResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 镜像的漏洞列表 **取值范围**: 取值0-2000个ImageVulInfo对象
	DataList       *[]ImageVulInfo `json:"data_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListLocalImageVulResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLocalImageVulResponse struct{}"
	}

	return strings.Join([]string{"ListLocalImageVulResponse", string(data)}, " ")
}
