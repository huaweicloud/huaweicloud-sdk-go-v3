package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsDataRepository 文件系统 obs 数据存储库
type ObsDataRepository struct {

	// obs 桶名称
	Bucket string `json:"bucket"`

	// obs 桶 endpoint
	Endpoint string `json:"endpoint"`
}

func (o ObsDataRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsDataRepository struct{}"
	}

	return strings.Join([]string{"ObsDataRepository", string(data)}, " ")
}
