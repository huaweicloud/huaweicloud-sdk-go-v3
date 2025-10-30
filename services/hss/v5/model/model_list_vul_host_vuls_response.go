package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostVulsResponse Response Object
type ListVulHostVulsResponse struct {

	// array
	DataList *[]VulUrgentResponseInfoDataList `json:"data_list,omitempty"`

	// **参数解释**： 总数 **取值范围**： 最小值0，最大值100000000总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVulHostVulsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostVulsResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostVulsResponse", string(data)}, " ")
}
