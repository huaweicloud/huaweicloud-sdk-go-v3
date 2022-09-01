package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecordCallbackConfigRequest struct {

	// 配置ID，在创建配置成功后返回
	Id string `json:"id" xml:"id"`

	Body *RecordCallbackConfigRequest `json:"body,omitempty" xml:"body"`
}

func (o UpdateRecordCallbackConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordCallbackConfigRequest", string(data)}, " ")
}
