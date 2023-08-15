package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeBackupEncryptStatusRequest Request Object
type DescribeBackupEncryptStatusRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DescribeBackupEncryptStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeBackupEncryptStatusRequest struct{}"
	}

	return strings.Join([]string{"DescribeBackupEncryptStatusRequest", string(data)}, " ")
}
