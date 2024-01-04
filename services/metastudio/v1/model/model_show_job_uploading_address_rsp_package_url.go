package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobUploadingAddressRspPackageUrl 整包上传任务的url。
type ShowJobUploadingAddressRspPackageUrl struct {

	// 上传的训练数据地址,用户需要将训练数据打成zip包后上传到该url。 > * 通过该obs地址上传时需要设置content-type为application/zip
	TrainingDataUploadingUrl *string `json:"training_data_uploading_url,omitempty"`
}

func (o ShowJobUploadingAddressRspPackageUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobUploadingAddressRspPackageUrl struct{}"
	}

	return strings.Join([]string{"ShowJobUploadingAddressRspPackageUrl", string(data)}, " ")
}
