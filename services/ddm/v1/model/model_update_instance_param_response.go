package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInstanceParamResponse struct {

	// 节点列表。
	NodeList *string `json:"nodeList,omitempty"`

	// 是否需要重启实例。
	NeedRestart *bool `json:"needRestart,omitempty"`

	// 任务id。
	JobId *string `json:"jobId,omitempty"`

	// 参数组id。
	ConfigId *string `json:"configId,omitempty"`

	// 参数组名称。
	ConfigName     *string `json:"configName,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceParamResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceParamResponse", string(data)}, " ")
}
