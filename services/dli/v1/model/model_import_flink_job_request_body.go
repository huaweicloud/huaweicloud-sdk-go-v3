package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobRequestBody 导入作业的请求参数。
type ImportFlinkJobRequestBody struct {

	// OBS上导入作业zip文件路径，支持填写目录，导入目录下所有zip文件。
	ZipFile string `json:"zip_file"`

	// 若导入作业中存在与服务已有作业同名情况，是否将服务中已有作业覆盖。
	IsCover *bool `json:"is_cover,omitempty"`
}

func (o ImportFlinkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobRequestBody struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobRequestBody", string(data)}, " ")
}
