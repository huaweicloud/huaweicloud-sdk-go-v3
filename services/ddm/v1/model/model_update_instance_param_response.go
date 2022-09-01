package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInstanceParamResponse struct {

	// 节点列表。
	NodeList *string `json:"nodeList,omitempty" xml:"nodeList"`

	// 是否需要重启实例。
	NeedRestart *bool `json:"needRestart,omitempty" xml:"needRestart"`

	// 任务id。
	JobId *string `json:"jobId,omitempty" xml:"jobId"`

	// 参数组id。
	ConfigId *string `json:"configId,omitempty" xml:"configId"`

	// 参数组名称。
	ConfigName     *string `json:"configName,omitempty" xml:"configName"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceParamResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceParamResponse", string(data)}, " ")
}
