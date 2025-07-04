package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShareFeatureGatesResponse Response Object
type ShowShareFeatureGatesResponse struct {

	// 是否支持体验馆
	EnableExperience *bool `json:"enable_experience,omitempty"`

	// 是否支持对接hss服务
	EnableHssService *bool `json:"enable_hss_service,omitempty"`

	// 是否支持镜像扫描
	EnableImageScan *bool `json:"enable_image_scan,omitempty"`

	// 是否支持国密场景
	EnableSm3 *bool `json:"enable_sm3,omitempty"`

	// 是否支持镜像同步
	EnableImageSync *bool `json:"enable_image_sync,omitempty"`

	// 是否支持对接cci服务
	EnableCciService *bool `json:"enable_cci_service,omitempty"`

	// 是否支持镜像标签
	EnableImageLabel *bool `json:"enable_image_label,omitempty"`

	// 是否支持流水线服务
	EnablePipeline *bool `json:"enable_pipeline,omitempty"`

	// 是否支持list v3接口
	EnableListV3   *bool `json:"enable_list_v3,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowShareFeatureGatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareFeatureGatesResponse struct{}"
	}

	return strings.Join([]string{"ShowShareFeatureGatesResponse", string(data)}, " ")
}
