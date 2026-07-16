package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEndpointsResp 远程接入训练作业时需要的相关配置。
type JobEndpointsResp struct {
	Ssh *SshResp `json:"ssh,omitempty"`

	JupyterLab *JupyterLab `json:"jupyter_lab,omitempty"`

	Tensorboard *Tensorboard `json:"tensorboard,omitempty"`

	MindstudioInsight *MindStudioInsight `json:"mindstudio_insight,omitempty"`
}

func (o JobEndpointsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEndpointsResp struct{}"
	}

	return strings.Join([]string{"JobEndpointsResp", string(data)}, " ")
}
