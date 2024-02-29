package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateEntitiesRequest Request Object
type CreateOrUpdateEntitiesRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)
	Instance string `json:"instance"`

	Body *DataEntityWithExtInfo `json:"body,omitempty"`
}

func (o CreateOrUpdateEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateEntitiesRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateEntitiesRequest", string(data)}, " ")
}
