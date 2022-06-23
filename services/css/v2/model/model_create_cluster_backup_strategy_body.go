package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 开启自动创建快照策略。  当backupStrategy参数配置不为空时，才会开启自动创建快照策略。
type CreateClusterBackupStrategyBody struct {

	// 每天自动创建快照的时间点。只支持整点，后面需加上时区，格式为“HH:mm z”，“HH:mm”表示整点时间，“z”表示时区。比如“00:00 GMT+08:00”、“01:00 GMT+08:00”等。
	Period string `json:"period"`

	// 自动创建的快照的前缀。
	Prefix string `json:"prefix"`

	// 自动创建快照的保留天数。取值范围：1-90。
	Keepday int32 `json:"keepday"`

	// 备份使用的OBS桶名称，如果桶已经存放快照数据了，不可更改。
	Bucket *string `json:"bucket,omitempty"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"basePath,omitempty"`

	// 访问OBS使用的IAM委托名称。   >如果bucket、basePath和agency三个参数同时为空，则系统会自动创建OBS桶和IAM代理，否则使用配置的参数值。
	Agency *string `json:"agency,omitempty"`
}

func (o CreateClusterBackupStrategyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterBackupStrategyBody struct{}"
	}

	return strings.Join([]string{"CreateClusterBackupStrategyBody", string(data)}, " ")
}
