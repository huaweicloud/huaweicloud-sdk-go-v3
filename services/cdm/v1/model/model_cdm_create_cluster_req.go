package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmCreateClusterReq struct {
	Cluster *CdmCreateClusterReqCluster `json:"cluster" xml:"cluster"`

	// 选择是否开启消息通知。开启后，支持配置5个手机号码或邮箱，作业（目前仅支持表/文件迁移的作业）失败时、EIP异常时会发送短信或邮件通知用户
	AutoRemind *bool `json:"auto_remind,omitempty" xml:"auto_remind"`

	// 接收消息通知的手机号码
	PhoneNum *string `json:"phone_num,omitempty" xml:"phone_num"`

	// 接收消息通知的邮箱
	Email *string `json:"email,omitempty" xml:"email"`
}

func (o CdmCreateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmCreateClusterReq struct{}"
	}

	return strings.Join([]string{"CdmCreateClusterReq", string(data)}, " ")
}
