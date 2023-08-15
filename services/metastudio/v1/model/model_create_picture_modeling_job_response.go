package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePictureModelingJobResponse Response Object
type CreatePictureModelingJobResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人资产ID。
	ModelAssetId   *string `json:"model_asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePictureModelingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePictureModelingJobResponse struct{}"
	}

	return strings.Join([]string{"CreatePictureModelingJobResponse", string(data)}, " ")
}
