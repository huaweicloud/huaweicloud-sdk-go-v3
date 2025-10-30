package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageNonCompliantAppResponse Response Object
type ListImageNonCompliantAppResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 镜像的不合规软件信息 **取值范围**: 最小值0，最大值10241
	DataList       *[]ImageNonCompliantAppInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListImageNonCompliantAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageNonCompliantAppResponse struct{}"
	}

	return strings.Join([]string{"ListImageNonCompliantAppResponse", string(data)}, " ")
}
