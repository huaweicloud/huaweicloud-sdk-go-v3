package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartManualSecurityCheckRequestInfo 手动体检的配置信息
type StartManualSecurityCheckRequestInfo struct {

	// 体检内容，取值包含：asset,vul,baseline,event
	Content *[]string `json:"content,omitempty"`

	// 已选服务器ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 服务器是否选择全部，全选跟查询条件无关
	OperateAll *bool `json:"operate_all,omitempty"`
}

func (o StartManualSecurityCheckRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartManualSecurityCheckRequestInfo struct{}"
	}

	return strings.Join([]string{"StartManualSecurityCheckRequestInfo", string(data)}, " ")
}
