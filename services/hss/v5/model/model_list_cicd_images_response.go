package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCicdImagesResponse Response Object
type ListCicdImagesResponse struct {

	// **参数解释**: cicd镜像总数 **取值范围**: 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: cicd镜像列表 **取值范围**: 最小值0，最大值200
	DataList       *[]CicdImageResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCicdImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCicdImagesResponse struct{}"
	}

	return strings.Join([]string{"ListCicdImagesResponse", string(data)}, " ")
}
