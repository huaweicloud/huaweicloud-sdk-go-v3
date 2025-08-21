package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataMapLineageRequest Request Object
type ImportDataMapLineageRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	Body *ImportDataMapLineageRequestBody `json:"body,omitempty"`
}

func (o ImportDataMapLineageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataMapLineageRequest struct{}"
	}

	return strings.Join([]string{"ImportDataMapLineageRequest", string(data)}, " ")
}
