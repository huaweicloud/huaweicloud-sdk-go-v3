package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicEntitiesRequest Request Object
type ListLogicEntitiesRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	// 主题目录资产guid，获取方法请参见[数据资产guid](dataartsstudio_02_0351.xml)。
	Guid string `json:"guid"`
}

func (o ListLogicEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicEntitiesRequest struct{}"
	}

	return strings.Join([]string{"ListLogicEntitiesRequest", string(data)}, " ")
}
