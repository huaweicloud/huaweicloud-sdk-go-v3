package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntityDetailsRequest Request Object
type ListEntityDetailsRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	// 批量获取资产信息请求体。
	Body *[]string `json:"body,omitempty"`
}

func (o ListEntityDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntityDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListEntityDetailsRequest", string(data)}, " ")
}
