package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryTaskCompletionRequest Request Object
type ListFactoryTaskCompletionRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 查询任务的类型，默认为all，查询所有任务。 类型有：Dummy、CDM Job、MRS Hive SQL、MRS Spark SQL、MRS Impala SQL、DLI SQL、DLI Spark、Python、DWS SQL、Shell、MRS ClickHouse、MRS HetuEngine
	Type *string `json:"type,omitempty"`
}

func (o ListFactoryTaskCompletionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryTaskCompletionRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryTaskCompletionRequest", string(data)}, " ")
}
