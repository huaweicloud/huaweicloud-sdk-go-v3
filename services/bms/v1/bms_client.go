package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bms/v1/model"
)

type BmsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewBmsClient(hcClient *httpclient.HcHttpClient) *BmsClient {
	return &BmsClient{HcClient: hcClient}
}

func BmsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddServerNics 裸金属服务器绑定弹性网卡
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) AddServerNics(request *model.AddServerNicsRequest) (*model.AddServerNicsResponse, error) {
	requestDef := GenReqDefForAddServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServerNicsResponse), nil
	}
}

// AddServerNicsInvoker 裸金属服务器绑定弹性网卡
func (c *BmsClient) AddServerNicsInvoker(request *model.AddServerNicsRequest) *AddServerNicsInvoker {
	requestDef := GenReqDefForAddServerNics()
	return &AddServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachBaremetalServerVolume 裸金属服务器挂载云硬盘
//
// 裸金属服务器创建成功后，如果发现磁盘不够用或者当前磁盘不满足要求，可以将已有云硬盘挂载给裸金属服务器，作为数据盘使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) AttachBaremetalServerVolume(request *model.AttachBaremetalServerVolumeRequest) (*model.AttachBaremetalServerVolumeResponse, error) {
	requestDef := GenReqDefForAttachBaremetalServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachBaremetalServerVolumeResponse), nil
	}
}

// AttachBaremetalServerVolumeInvoker 裸金属服务器挂载云硬盘
func (c *BmsClient) AttachBaremetalServerVolumeInvoker(request *model.AttachBaremetalServerVolumeRequest) *AttachBaremetalServerVolumeInvoker {
	requestDef := GenReqDefForAttachBaremetalServerVolume()
	return &AttachBaremetalServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateBaremetalServerTags 批量添加裸金属服务器标签
//
// - 为指定裸金属服务器批量添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) BatchCreateBaremetalServerTags(request *model.BatchCreateBaremetalServerTagsRequest) (*model.BatchCreateBaremetalServerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateBaremetalServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateBaremetalServerTagsResponse), nil
	}
}

// BatchCreateBaremetalServerTagsInvoker 批量添加裸金属服务器标签
func (c *BmsClient) BatchCreateBaremetalServerTagsInvoker(request *model.BatchCreateBaremetalServerTagsRequest) *BatchCreateBaremetalServerTagsInvoker {
	requestDef := GenReqDefForBatchCreateBaremetalServerTags()
	return &BatchCreateBaremetalServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteBaremetalServerTags 批量删除裸金属服务器标签
//
// - 为指定云服务器批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) BatchDeleteBaremetalServerTags(request *model.BatchDeleteBaremetalServerTagsRequest) (*model.BatchDeleteBaremetalServerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteBaremetalServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteBaremetalServerTagsResponse), nil
	}
}

// BatchDeleteBaremetalServerTagsInvoker 批量删除裸金属服务器标签
func (c *BmsClient) BatchDeleteBaremetalServerTagsInvoker(request *model.BatchDeleteBaremetalServerTagsRequest) *BatchDeleteBaremetalServerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteBaremetalServerTags()
	return &BatchDeleteBaremetalServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebootBaremetalServers 重启裸金属服务器
//
// 根据给定的裸金属服务器ID列表，批量重启裸金属服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) BatchRebootBaremetalServers(request *model.BatchRebootBaremetalServersRequest) (*model.BatchRebootBaremetalServersResponse, error) {
	requestDef := GenReqDefForBatchRebootBaremetalServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootBaremetalServersResponse), nil
	}
}

// BatchRebootBaremetalServersInvoker 重启裸金属服务器
func (c *BmsClient) BatchRebootBaremetalServersInvoker(request *model.BatchRebootBaremetalServersRequest) *BatchRebootBaremetalServersInvoker {
	requestDef := GenReqDefForBatchRebootBaremetalServers()
	return &BatchRebootBaremetalServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartBaremetalServers 启动裸金属服务器
//
// 根据给定的裸金属服务器ID列表，批量启动裸金属服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) BatchStartBaremetalServers(request *model.BatchStartBaremetalServersRequest) (*model.BatchStartBaremetalServersResponse, error) {
	requestDef := GenReqDefForBatchStartBaremetalServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartBaremetalServersResponse), nil
	}
}

// BatchStartBaremetalServersInvoker 启动裸金属服务器
func (c *BmsClient) BatchStartBaremetalServersInvoker(request *model.BatchStartBaremetalServersRequest) *BatchStartBaremetalServersInvoker {
	requestDef := GenReqDefForBatchStartBaremetalServers()
	return &BatchStartBaremetalServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopBaremetalServers 关闭裸金属服务器
//
// 根据给定的裸金属服务器ID列表，批量关闭裸金属服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) BatchStopBaremetalServers(request *model.BatchStopBaremetalServersRequest) (*model.BatchStopBaremetalServersResponse, error) {
	requestDef := GenReqDefForBatchStopBaremetalServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopBaremetalServersResponse), nil
	}
}

// BatchStopBaremetalServersInvoker 关闭裸金属服务器
func (c *BmsClient) BatchStopBaremetalServersInvoker(request *model.BatchStopBaremetalServersRequest) *BatchStopBaremetalServersInvoker {
	requestDef := GenReqDefForBatchStopBaremetalServers()
	return &BatchStopBaremetalServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeBaremetalServerName 修改裸金属服务器名称
//
// 修改裸金属服务器名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ChangeBaremetalServerName(request *model.ChangeBaremetalServerNameRequest) (*model.ChangeBaremetalServerNameResponse, error) {
	requestDef := GenReqDefForChangeBaremetalServerName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeBaremetalServerNameResponse), nil
	}
}

// ChangeBaremetalServerNameInvoker 修改裸金属服务器名称
func (c *BmsClient) ChangeBaremetalServerNameInvoker(request *model.ChangeBaremetalServerNameRequest) *ChangeBaremetalServerNameInvoker {
	requestDef := GenReqDefForChangeBaremetalServerName()
	return &ChangeBaremetalServerNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeBaremetalServerOs 切换裸金属服务器的操作系统
//
// 切换裸金属服务器的操作系统。切换操作系统支持密码或者密钥注入，该接口支持企业项目细粒度权限的校验，具体细粒度请参见 bms:servers:changeOS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ChangeBaremetalServerOs(request *model.ChangeBaremetalServerOsRequest) (*model.ChangeBaremetalServerOsResponse, error) {
	requestDef := GenReqDefForChangeBaremetalServerOs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeBaremetalServerOsResponse), nil
	}
}

// ChangeBaremetalServerOsInvoker 切换裸金属服务器的操作系统
func (c *BmsClient) ChangeBaremetalServerOsInvoker(request *model.ChangeBaremetalServerOsRequest) *ChangeBaremetalServerOsInvoker {
	requestDef := GenReqDefForChangeBaremetalServerOs()
	return &ChangeBaremetalServerOsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBareMetalServers 创建裸金属服务器
//
// 创建一台或多台裸金属服务器,裸金属服务器的登录鉴权方式包括两种：密钥对、密码。为安全起见，推荐使用密钥对方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) CreateBareMetalServers(request *model.CreateBareMetalServersRequest) (*model.CreateBareMetalServersResponse, error) {
	requestDef := GenReqDefForCreateBareMetalServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBareMetalServersResponse), nil
	}
}

// CreateBareMetalServersInvoker 创建裸金属服务器
func (c *BmsClient) CreateBareMetalServersInvoker(request *model.CreateBareMetalServersRequest) *CreateBareMetalServersInvoker {
	requestDef := GenReqDefForCreateBareMetalServers()
	return &CreateBareMetalServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBaremetalServer 删除裸金属服务器裸金属服务器物理机
//
// 删除裸金属服务器裸金属服务器物理机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) DeleteBaremetalServer(request *model.DeleteBaremetalServerRequest) (*model.DeleteBaremetalServerResponse, error) {
	requestDef := GenReqDefForDeleteBaremetalServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBaremetalServerResponse), nil
	}
}

// DeleteBaremetalServerInvoker 删除裸金属服务器裸金属服务器物理机
func (c *BmsClient) DeleteBaremetalServerInvoker(request *model.DeleteBaremetalServerRequest) *DeleteBaremetalServerInvoker {
	requestDef := GenReqDefForDeleteBaremetalServer()
	return &DeleteBaremetalServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerNics 裸金属服务器解绑弹性网卡
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) DeleteServerNics(request *model.DeleteServerNicsRequest) (*model.DeleteServerNicsResponse, error) {
	requestDef := GenReqDefForDeleteServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerNicsResponse), nil
	}
}

// DeleteServerNicsInvoker 裸金属服务器解绑弹性网卡
func (c *BmsClient) DeleteServerNicsInvoker(request *model.DeleteServerNicsRequest) *DeleteServerNicsInvoker {
	requestDef := GenReqDefForDeleteServerNics()
	return &DeleteServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWindowsBareMetalServerPassword Windows裸金属服务器清除密码
//
// 清除Windows裸金属服务器初始安装时系统生成的密码记录。清除密码后，不影响裸金属服务器密码登录功能，但不能再使用获取密码功能来查询该裸金属服务器密码。如果裸金属服务器是通过私有镜像创建的，请确保已安装Cloudbase-init。公共镜像默认已安装该软件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) DeleteWindowsBareMetalServerPassword(request *model.DeleteWindowsBareMetalServerPasswordRequest) (*model.DeleteWindowsBareMetalServerPasswordResponse, error) {
	requestDef := GenReqDefForDeleteWindowsBareMetalServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWindowsBareMetalServerPasswordResponse), nil
	}
}

// DeleteWindowsBareMetalServerPasswordInvoker Windows裸金属服务器清除密码
func (c *BmsClient) DeleteWindowsBareMetalServerPasswordInvoker(request *model.DeleteWindowsBareMetalServerPasswordRequest) *DeleteWindowsBareMetalServerPasswordInvoker {
	requestDef := GenReqDefForDeleteWindowsBareMetalServerPassword()
	return &DeleteWindowsBareMetalServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachBaremetalServerVolume 裸金属服务器卸载云磁盘
//
// 将挂载至裸金属服务器中的磁盘卸载；对于挂载在系统盘盘位（也就是“/dev/sda”挂载点）上的磁盘，不允许执行卸载操作；对于挂载在数据盘盘位（非“/dev/sda”挂载点）上的磁盘，支持离线卸载和在线卸载（裸金属服务器处于“运行中”状态）磁盘
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) DetachBaremetalServerVolume(request *model.DetachBaremetalServerVolumeRequest) (*model.DetachBaremetalServerVolumeResponse, error) {
	requestDef := GenReqDefForDetachBaremetalServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachBaremetalServerVolumeResponse), nil
	}
}

// DetachBaremetalServerVolumeInvoker 裸金属服务器卸载云磁盘
func (c *BmsClient) DetachBaremetalServerVolumeInvoker(request *model.DetachBaremetalServerVolumeRequest) *DetachBaremetalServerVolumeInvoker {
	requestDef := GenReqDefForDetachBaremetalServerVolume()
	return &DetachBaremetalServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBareMetalServerDetails 查询裸金属服务器详情
//
// 获取裸金属服务器详细信息，该接口支持查询裸金属服务器的计费方式，以及是否被冻结
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ListBareMetalServerDetails(request *model.ListBareMetalServerDetailsRequest) (*model.ListBareMetalServerDetailsResponse, error) {
	requestDef := GenReqDefForListBareMetalServerDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBareMetalServerDetailsResponse), nil
	}
}

// ListBareMetalServerDetailsInvoker 查询裸金属服务器详情
func (c *BmsClient) ListBareMetalServerDetailsInvoker(request *model.ListBareMetalServerDetailsRequest) *ListBareMetalServerDetailsInvoker {
	requestDef := GenReqDefForListBareMetalServerDetails()
	return &ListBareMetalServerDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBareMetalServers 查询裸金属服务器详情列表
//
// 用户根据设置的请求条件筛选裸金属服务器，并获取裸金属服务器的详细信息。该接口支持查询裸金属服务器计费方式，以及是否被冻结。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ListBareMetalServers(request *model.ListBareMetalServersRequest) (*model.ListBareMetalServersResponse, error) {
	requestDef := GenReqDefForListBareMetalServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBareMetalServersResponse), nil
	}
}

// ListBareMetalServersInvoker 查询裸金属服务器详情列表
func (c *BmsClient) ListBareMetalServersInvoker(request *model.ListBareMetalServersRequest) *ListBareMetalServersInvoker {
	requestDef := GenReqDefForListBareMetalServers()
	return &ListBareMetalServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBareMetalServersDetail 查询裸金属服务器列表
//
// 用户根据设置的请求条件筛选裸金属服务器，并获取裸金属服务器的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ListBareMetalServersDetail(request *model.ListBareMetalServersDetailRequest) (*model.ListBareMetalServersDetailResponse, error) {
	requestDef := GenReqDefForListBareMetalServersDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBareMetalServersDetailResponse), nil
	}
}

// ListBareMetalServersDetailInvoker 查询裸金属服务器列表
func (c *BmsClient) ListBareMetalServersDetailInvoker(request *model.ListBareMetalServersDetailRequest) *ListBareMetalServersDetailInvoker {
	requestDef := GenReqDefForListBareMetalServersDetail()
	return &ListBareMetalServersDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBaremetalFlavorDetailExtends 查询规格详情和规格扩展信息列表
//
// 查询裸金属服务器的规格详情和规格的扩展信息。您可以调用此接口查询“baremetal:extBootType”参数取值，以确认某个规格是否支持快速发放
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ListBaremetalFlavorDetailExtends(request *model.ListBaremetalFlavorDetailExtendsRequest) (*model.ListBaremetalFlavorDetailExtendsResponse, error) {
	requestDef := GenReqDefForListBaremetalFlavorDetailExtends()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBaremetalFlavorDetailExtendsResponse), nil
	}
}

// ListBaremetalFlavorDetailExtendsInvoker 查询规格详情和规格扩展信息列表
func (c *BmsClient) ListBaremetalFlavorDetailExtendsInvoker(request *model.ListBaremetalFlavorDetailExtendsRequest) *ListBaremetalFlavorDetailExtendsInvoker {
	requestDef := GenReqDefForListBaremetalFlavorDetailExtends()
	return &ListBaremetalFlavorDetailExtendsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallBaremetalServerOs 重装裸金属服务器操作系统
//
// 重装裸金属服务器的操作系统。快速发放裸金属服务器支持裸金属服务器数据盘不变的情况下，使用原镜像重装系统盘。重装操作系统支持密码或者密钥注入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ReinstallBaremetalServerOs(request *model.ReinstallBaremetalServerOsRequest) (*model.ReinstallBaremetalServerOsResponse, error) {
	requestDef := GenReqDefForReinstallBaremetalServerOs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallBaremetalServerOsResponse), nil
	}
}

// ReinstallBaremetalServerOsInvoker 重装裸金属服务器操作系统
func (c *BmsClient) ReinstallBaremetalServerOsInvoker(request *model.ReinstallBaremetalServerOsRequest) *ReinstallBaremetalServerOsInvoker {
	requestDef := GenReqDefForReinstallBaremetalServerOs()
	return &ReinstallBaremetalServerOsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwdOneClick 一键重置裸金属服务器密码
//
// 在裸金属服务器支持一键重置密码功能的前提下，重置裸金属服务器管理帐号（root用户或Administrator用户）的密码。可以通过6.10.1-查询是否支持一键重置密码API查询是否支持一键重置密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ResetPwdOneClick(request *model.ResetPwdOneClickRequest) (*model.ResetPwdOneClickResponse, error) {
	requestDef := GenReqDefForResetPwdOneClick()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdOneClickResponse), nil
	}
}

// ResetPwdOneClickInvoker 一键重置裸金属服务器密码
func (c *BmsClient) ResetPwdOneClickInvoker(request *model.ResetPwdOneClickRequest) *ResetPwdOneClickInvoker {
	requestDef := GenReqDefForResetPwdOneClick()
	return &ResetPwdOneClickInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaremetalServerInterfaceAttachments 查询裸金属服务器网卡信息
//
// 查询裸金属服务器的网卡信息，比如网卡的IP地址、MAC地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowBaremetalServerInterfaceAttachments(request *model.ShowBaremetalServerInterfaceAttachmentsRequest) (*model.ShowBaremetalServerInterfaceAttachmentsResponse, error) {
	requestDef := GenReqDefForShowBaremetalServerInterfaceAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaremetalServerInterfaceAttachmentsResponse), nil
	}
}

// ShowBaremetalServerInterfaceAttachmentsInvoker 查询裸金属服务器网卡信息
func (c *BmsClient) ShowBaremetalServerInterfaceAttachmentsInvoker(request *model.ShowBaremetalServerInterfaceAttachmentsRequest) *ShowBaremetalServerInterfaceAttachmentsInvoker {
	requestDef := GenReqDefForShowBaremetalServerInterfaceAttachments()
	return &ShowBaremetalServerInterfaceAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaremetalServerTags 查询裸金属服务器标签
//
// - 查询指定云服务器的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowBaremetalServerTags(request *model.ShowBaremetalServerTagsRequest) (*model.ShowBaremetalServerTagsResponse, error) {
	requestDef := GenReqDefForShowBaremetalServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaremetalServerTagsResponse), nil
	}
}

// ShowBaremetalServerTagsInvoker 查询裸金属服务器标签
func (c *BmsClient) ShowBaremetalServerTagsInvoker(request *model.ShowBaremetalServerTagsRequest) *ShowBaremetalServerTagsInvoker {
	requestDef := GenReqDefForShowBaremetalServerTags()
	return &ShowBaremetalServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaremetalServerVolumeInfo 查询裸金属服务器挂载的云硬盘信息
//
// 查询裸金属服务器挂载的磁盘信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowBaremetalServerVolumeInfo(request *model.ShowBaremetalServerVolumeInfoRequest) (*model.ShowBaremetalServerVolumeInfoResponse, error) {
	requestDef := GenReqDefForShowBaremetalServerVolumeInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaremetalServerVolumeInfoResponse), nil
	}
}

// ShowBaremetalServerVolumeInfoInvoker 查询裸金属服务器挂载的云硬盘信息
func (c *BmsClient) ShowBaremetalServerVolumeInfoInvoker(request *model.ShowBaremetalServerVolumeInfoRequest) *ShowBaremetalServerVolumeInfoInvoker {
	requestDef := GenReqDefForShowBaremetalServerVolumeInfo()
	return &ShowBaremetalServerVolumeInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResetPwd 查询是否支持一键重置密码
//
// 查询是否支持一键重置密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowResetPwd(request *model.ShowResetPwdRequest) (*model.ShowResetPwdResponse, error) {
	requestDef := GenReqDefForShowResetPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResetPwdResponse), nil
	}
}

// ShowResetPwdInvoker 查询是否支持一键重置密码
func (c *BmsClient) ShowResetPwdInvoker(request *model.ShowResetPwdRequest) *ShowResetPwdInvoker {
	requestDef := GenReqDefForShowResetPwd()
	return &ShowResetPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerRemoteConsole 获取裸金属服务器远程登录地址
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowServerRemoteConsole(request *model.ShowServerRemoteConsoleRequest) (*model.ShowServerRemoteConsoleResponse, error) {
	requestDef := GenReqDefForShowServerRemoteConsole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerRemoteConsoleResponse), nil
	}
}

// ShowServerRemoteConsoleInvoker 获取裸金属服务器远程登录地址
func (c *BmsClient) ShowServerRemoteConsoleInvoker(request *model.ShowServerRemoteConsoleRequest) *ShowServerRemoteConsoleInvoker {
	requestDef := GenReqDefForShowServerRemoteConsole()
	return &ShowServerRemoteConsoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantQuota 查询租户配额
//
// 查询该租户下，所有资源的配额信息，包括已使用配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowTenantQuota(request *model.ShowTenantQuotaRequest) (*model.ShowTenantQuotaResponse, error) {
	requestDef := GenReqDefForShowTenantQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantQuotaResponse), nil
	}
}

// ShowTenantQuotaInvoker 查询租户配额
func (c *BmsClient) ShowTenantQuotaInvoker(request *model.ShowTenantQuotaRequest) *ShowTenantQuotaInvoker {
	requestDef := GenReqDefForShowTenantQuota()
	return &ShowTenantQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWindowsBaremetalServerPwd Windows裸金属服务器获取密码
//
// 获取Windows裸金属服务器初始安装时系统生成的管理员帐户（Administrator帐户或Cloudbase-init设置的帐户）随机密码。如果裸金属服务器是通过私有镜像创建的，请确保已安装Cloudbase-init。公共镜像默认已安装该软件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowWindowsBaremetalServerPwd(request *model.ShowWindowsBaremetalServerPwdRequest) (*model.ShowWindowsBaremetalServerPwdResponse, error) {
	requestDef := GenReqDefForShowWindowsBaremetalServerPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWindowsBaremetalServerPwdResponse), nil
	}
}

// ShowWindowsBaremetalServerPwdInvoker Windows裸金属服务器获取密码
func (c *BmsClient) ShowWindowsBaremetalServerPwdInvoker(request *model.ShowWindowsBaremetalServerPwdRequest) *ShowWindowsBaremetalServerPwdInvoker {
	requestDef := GenReqDefForShowWindowsBaremetalServerPwd()
	return &ShowWindowsBaremetalServerPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBaremetalServerInterfaceAttachments 修改裸金属服务器弹性网卡的属性
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) UpdateBaremetalServerInterfaceAttachments(request *model.UpdateBaremetalServerInterfaceAttachmentsRequest) (*model.UpdateBaremetalServerInterfaceAttachmentsResponse, error) {
	requestDef := GenReqDefForUpdateBaremetalServerInterfaceAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBaremetalServerInterfaceAttachmentsResponse), nil
	}
}

// UpdateBaremetalServerInterfaceAttachmentsInvoker 修改裸金属服务器弹性网卡的属性
func (c *BmsClient) UpdateBaremetalServerInterfaceAttachmentsInvoker(request *model.UpdateBaremetalServerInterfaceAttachmentsRequest) *UpdateBaremetalServerInterfaceAttachmentsInvoker {
	requestDef := GenReqDefForUpdateBaremetalServerInterfaceAttachments()
	return &UpdateBaremetalServerInterfaceAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBaremetalServerMetadata 更新裸金属服务器元数据
//
// 更新裸金属服务器元数据。如果元数据中没有待更新字段，则自动添加该字段。如果元数据中已存在待更新字段，则直接更新字段值；如果元数据中的字段不再请求参数中，则保持不变
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) UpdateBaremetalServerMetadata(request *model.UpdateBaremetalServerMetadataRequest) (*model.UpdateBaremetalServerMetadataResponse, error) {
	requestDef := GenReqDefForUpdateBaremetalServerMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBaremetalServerMetadataResponse), nil
	}
}

// UpdateBaremetalServerMetadataInvoker 更新裸金属服务器元数据
func (c *BmsClient) UpdateBaremetalServerMetadataInvoker(request *model.UpdateBaremetalServerMetadataRequest) *UpdateBaremetalServerMetadataInvoker {
	requestDef := GenReqDefForUpdateBaremetalServerMetadata()
	return &UpdateBaremetalServerMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpecifiedVersion 查询指定API版本信息
//
// 查询裸金属服务指定接口版本的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowSpecifiedVersion(request *model.ShowSpecifiedVersionRequest) (*model.ShowSpecifiedVersionResponse, error) {
	requestDef := GenReqDefForShowSpecifiedVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpecifiedVersionResponse), nil
	}
}

// ShowSpecifiedVersionInvoker 查询指定API版本信息
func (c *BmsClient) ShowSpecifiedVersionInvoker(request *model.ShowSpecifiedVersionRequest) *ShowSpecifiedVersionInvoker {
	requestDef := GenReqDefForShowSpecifiedVersion()
	return &ShowSpecifiedVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobInfos 查询Job状态
//
// 查询Job的执行状态。对于创建裸金属服务器物理机、挂卸卷等异步API，命令下发后，会返回job_id，通过job_id可以查询任务的执行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BmsClient) ShowJobInfos(request *model.ShowJobInfosRequest) (*model.ShowJobInfosResponse, error) {
	requestDef := GenReqDefForShowJobInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobInfosResponse), nil
	}
}

// ShowJobInfosInvoker 查询Job状态
func (c *BmsClient) ShowJobInfosInvoker(request *model.ShowJobInfosRequest) *ShowJobInfosInvoker {
	requestDef := GenReqDefForShowJobInfos()
	return &ShowJobInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
