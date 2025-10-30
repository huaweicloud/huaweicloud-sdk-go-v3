package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageBuildCommandRisksImagesResponse Response Object
type ListImageBuildCommandRisksImagesResponse struct {

	// **参数解释** 镜像构建指令影响的镜像数据总量 **取值范围** 取值0-2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 镜像构建指令影响的镜像列表 **取值范围**: 取值0-200
	DataList       *[]BuildCommandRisksImageResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListImageBuildCommandRisksImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageBuildCommandRisksImagesResponse struct{}"
	}

	return strings.Join([]string{"ListImageBuildCommandRisksImagesResponse", string(data)}, " ")
}
