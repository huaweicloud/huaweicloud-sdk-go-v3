package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseLevelDatabaseResponse Response Object
type ShowDatabaseLevelDatabaseResponse struct {

	// 库信息列表
	Databases *[]DatabaseBackupInfo `json:"databases,omitempty"`

	// 可恢复库的个数
	DatabaseLimit *int32 `json:"database_limit,omitempty"`

	// obs桶名
	BucketName     *string `json:"bucket_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDatabaseLevelDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseLevelDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseLevelDatabaseResponse", string(data)}, " ")
}
