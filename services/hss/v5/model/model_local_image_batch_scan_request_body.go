package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LocalImageBatchScanRequestBody struct {

	// 是否扫描全部本地镜像，true扫描所有，非true则不是
	OperateAll *bool `json:"operate_all,omitempty"`

	// 需要扫描的本地镜像列表
	ImageInfoList *[]LocalImageInfo `json:"image_info_list,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像类型，包含如下:   - other_image : 其他镜像   - swr_image : swr镜像仓
	LocalImageType *string `json:"local_image_type,omitempty"`

	// 扫描状态，包含如下:   - unscan : 未扫描   - success : 扫描完成   - scanning : 扫描中   - failed : 扫描失败   - download_failed : 下载失败   - image_oversized : 镜像超大
	ScanStatus *string `json:"scan_status,omitempty"`

	// 镜像大小
	ImageSize *int64 `json:"image_size,omitempty"`

	// 创建时间开始日期，时间单位：毫秒（ms）
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// 创建时间结束日期，时间单位：毫秒（ms）
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// 最近一次扫描完成时间开始日期，时间单位：毫秒（ms）
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// 最近一次扫描完成时间结束日期，时间单位：毫秒（ms）
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`
}

func (o LocalImageBatchScanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalImageBatchScanRequestBody struct{}"
	}

	return strings.Join([]string{"LocalImageBatchScanRequestBody", string(data)}, " ")
}
