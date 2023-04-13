package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePictureModelingByUrlJobResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人资产ID。
	ModelAssetId   *string `json:"model_asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePictureModelingByUrlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePictureModelingByUrlJobResponse struct{}"
	}

	return strings.Join([]string{"CreatePictureModelingByUrlJobResponse", string(data)}, " ")
}
