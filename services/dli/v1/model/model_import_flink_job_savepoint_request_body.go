package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobSavepointRequestBody 导入savepoint接口的请求体
type ImportFlinkJobSavepointRequestBody struct {

	// Flink作业的id
	JobId int64 `json:"job_id"`

	// Savepoint路径。需指定到_metaData文件的上级目录 例：\"obs://bucket_name/file_name/\"
	SavepointPath string `json:"savepoint_path"`
}

func (o ImportFlinkJobSavepointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobSavepointRequestBody struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobSavepointRequestBody", string(data)}, " ")
}
