package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEntityRequest Request Object
type DeleteEntityRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	// 资产的guid，获取方法请参见[数据开发作业ID](dataartsstudio_02_0351.xml)。
	Guid string `json:"guid"`
}

func (o DeleteEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEntityRequest struct{}"
	}

	return strings.Join([]string{"DeleteEntityRequest", string(data)}, " ")
}
