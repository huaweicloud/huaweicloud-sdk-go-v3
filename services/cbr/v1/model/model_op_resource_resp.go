package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpResourceResp struct {
	ExtraInfo *OpResourceExtraInfoResp `json:"extra_info,omitempty"`

	// **参数解释：** 待备份资源ID，用于标识需要备份的资源。 [获取方法请参见\"[查询存储库绑定资源信息](https://support.huaweicloud.com/api-cbr/ShowVault.html)\"，\"[查询资源可保护性](https://support.huaweicloud.com/api-cbr/ShowProtectable.html)\"。](tag:hws) [获取方法请参见\"[查询存储库绑定资源信息](https://support.huaweicloud.com/intl/zh-cn/api-cbr/ShowVault.html)\"，\"[查询资源可保护性](https://support.huaweicloud.com/intl/zh-cn/api-cbr/ShowProtectable.html)\"。](tag:hws_hk) **取值范围：** 格式为UUID
	Id string `json:"id"`

	// **参数解释：** 待备份资源名称，用于标识待备份的资源名称。 **取值范围：** 长度为0~255个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 待备份资源的类型，用于指定待备份的资源类型，不同类型对应不同的云服务资源。 **取值范围：** [OS::Nova::Server：云服务器； OS::Cinder::Volume：云硬盘；OS::Ironic::BareMetalServer：裸金属服务器；OS::Sfs::Turbo：SFS Turbo文件系统；OS::Workspace::DesktopV2：云桌面。](tag:hws,hws_hk) [OS::Nova::Server：云服务器；OS::Cinder::Volume：云硬盘；OS::Sfs::Turbo：SFS Turbo文件系统。](tag:hk_g42,sbc,dt) [OS::Nova::Server：云服务器；OS::Cinder::Volume：云硬盘；OS::Ironic::BareMetalServer：裸金属服务器；OS::Sfs::Turbo：SFS Turbo文件系统。](tag:fcs_vm,ctc,ocb,tm) [OS::Nova::Server：云服务器；OS::Cinder::Volume：云硬盘。](tag:tlf,cmcc,hcso_dt) [OS::Nova::Server：云服务器；OS::Cinder::Volume：云硬盘；OS::Sfs::Turbo：SFS Turbo文件系统；OS::Workspace::DesktopV2：云桌面。](tag:g42)
	Type string `json:"type"`
}

func (o OpResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpResourceResp struct{}"
	}

	return strings.Join([]string{"OpResourceResp", string(data)}, " ")
}
