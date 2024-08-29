package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

type EcsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEcsClient(hcClient *httpclient.HcHttpClient) *EcsClient {
	return &EcsClient{HcClient: hcClient}
}

func EcsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddServerGroupMember 添加云服务器组成员
//
// 将云服务器加入云服务器组。添加成功后，如果该云服务器组是反亲和性策略的，则该云服务器与云服务器组中的其他成员尽量分散地创建在不同主机上。如果该云服务器时故障域类型的，则该云服务器会拥有故障域属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) AddServerGroupMember(request *model.AddServerGroupMemberRequest) (*model.AddServerGroupMemberResponse, error) {
	requestDef := GenReqDefForAddServerGroupMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServerGroupMemberResponse), nil
	}
}

// AddServerGroupMemberInvoker 添加云服务器组成员
func (c *EcsClient) AddServerGroupMemberInvoker(request *model.AddServerGroupMemberRequest) *AddServerGroupMemberInvoker {
	requestDef := GenReqDefForAddServerGroupMember()
	return &AddServerGroupMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateServerVirtualIp 云服务器网卡配置虚拟IP地址
//
// 虚拟IP地址用于为网卡提供第二个IP地址，同时支持与多个弹性云服务器的网卡绑定，从而实现多个弹性云服务器之间的高可用性。
//
// 该接口用于给云服务器网卡配置虚拟IP地址：
//
// - 当指定的IP地址是一个不存在的虚拟IP地址时，系统会创建该虚拟IP，并绑定至对应网卡。
//
// - 当指定的IP地址是一个已经创建好的私有IP时，系统会将指定的网卡和虚拟IP绑定。如果该IP的device_owner为空，则仅支持VPC内二三层通信；如果该IP的device_owner为neutron:VIP_PORT，则支持VPC内二三层通信、VPC之间对等连接访问，以及弹性公网IP、VPN、云专线等Internet接入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) AssociateServerVirtualIp(request *model.AssociateServerVirtualIpRequest) (*model.AssociateServerVirtualIpResponse, error) {
	requestDef := GenReqDefForAssociateServerVirtualIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateServerVirtualIpResponse), nil
	}
}

// AssociateServerVirtualIpInvoker 云服务器网卡配置虚拟IP地址
func (c *EcsClient) AssociateServerVirtualIpInvoker(request *model.AssociateServerVirtualIpRequest) *AssociateServerVirtualIpInvoker {
	requestDef := GenReqDefForAssociateServerVirtualIp()
	return &AssociateServerVirtualIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachServerVolume 弹性云服务器挂载磁盘
//
// 把磁盘挂载到弹性云服务器上。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) AttachServerVolume(request *model.AttachServerVolumeRequest) (*model.AttachServerVolumeResponse, error) {
	requestDef := GenReqDefForAttachServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachServerVolumeResponse), nil
	}
}

// AttachServerVolumeInvoker 弹性云服务器挂载磁盘
func (c *EcsClient) AttachServerVolumeInvoker(request *model.AttachServerVolumeRequest) *AttachServerVolumeInvoker {
	requestDef := GenReqDefForAttachServerVolume()
	return &AttachServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddServerNics 批量添加云服务器网卡
//
// 给云服务器添加一张或多张网卡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchAddServerNics(request *model.BatchAddServerNicsRequest) (*model.BatchAddServerNicsResponse, error) {
	requestDef := GenReqDefForBatchAddServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddServerNicsResponse), nil
	}
}

// BatchAddServerNicsInvoker 批量添加云服务器网卡
func (c *EcsClient) BatchAddServerNicsInvoker(request *model.BatchAddServerNicsRequest) *BatchAddServerNicsInvoker {
	requestDef := GenReqDefForBatchAddServerNics()
	return &BatchAddServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAttachSharableVolumes 批量挂载指定共享盘
//
// 将指定的共享磁盘一次性挂载到多个弹性云服务器，实现批量挂载。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchAttachSharableVolumes(request *model.BatchAttachSharableVolumesRequest) (*model.BatchAttachSharableVolumesResponse, error) {
	requestDef := GenReqDefForBatchAttachSharableVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAttachSharableVolumesResponse), nil
	}
}

// BatchAttachSharableVolumesInvoker 批量挂载指定共享盘
func (c *EcsClient) BatchAttachSharableVolumesInvoker(request *model.BatchAttachSharableVolumesRequest) *BatchAttachSharableVolumesInvoker {
	requestDef := GenReqDefForBatchAttachSharableVolumes()
	return &BatchAttachSharableVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateServerTags 批量添加云服务器标签
//
// - 为指定云服务器批量添加标签。
//
// - 标签管理服务TMS使用该接口批量管理云服务器的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchCreateServerTags(request *model.BatchCreateServerTagsRequest) (*model.BatchCreateServerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateServerTagsResponse), nil
	}
}

// BatchCreateServerTagsInvoker 批量添加云服务器标签
func (c *EcsClient) BatchCreateServerTagsInvoker(request *model.BatchCreateServerTagsRequest) *BatchCreateServerTagsInvoker {
	requestDef := GenReqDefForBatchCreateServerTags()
	return &BatchCreateServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServerNics 批量删除云服务器网卡
//
// 卸载并删除云服务器中的一张或多张网卡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchDeleteServerNics(request *model.BatchDeleteServerNicsRequest) (*model.BatchDeleteServerNicsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServerNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerNicsResponse), nil
	}
}

// BatchDeleteServerNicsInvoker 批量删除云服务器网卡
func (c *EcsClient) BatchDeleteServerNicsInvoker(request *model.BatchDeleteServerNicsRequest) *BatchDeleteServerNicsInvoker {
	requestDef := GenReqDefForBatchDeleteServerNics()
	return &BatchDeleteServerNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServerTags 批量删除云服务器标签
//
// - 为指定云服务器批量删除标签。
//
// - 标签管理服务TMS使用该接口批量管理云服务器的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchDeleteServerTags(request *model.BatchDeleteServerTagsRequest) (*model.BatchDeleteServerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerTagsResponse), nil
	}
}

// BatchDeleteServerTagsInvoker 批量删除云服务器标签
func (c *EcsClient) BatchDeleteServerTagsInvoker(request *model.BatchDeleteServerTagsRequest) *BatchDeleteServerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteServerTags()
	return &BatchDeleteServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebootServers 批量重启云服务器
//
// 根据给定的云服务器ID列表，批量重启云服务器，一次最多可以重启1000台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchRebootServers(request *model.BatchRebootServersRequest) (*model.BatchRebootServersResponse, error) {
	requestDef := GenReqDefForBatchRebootServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootServersResponse), nil
	}
}

// BatchRebootServersInvoker 批量重启云服务器
func (c *EcsClient) BatchRebootServersInvoker(request *model.BatchRebootServersRequest) *BatchRebootServersInvoker {
	requestDef := GenReqDefForBatchRebootServers()
	return &BatchRebootServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResetServersPassword 批量重置弹性云服务器密码
//
// 批量重置弹性云服务器管理帐号（root用户或Administrator用户）的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchResetServersPassword(request *model.BatchResetServersPasswordRequest) (*model.BatchResetServersPasswordResponse, error) {
	requestDef := GenReqDefForBatchResetServersPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetServersPasswordResponse), nil
	}
}

// BatchResetServersPasswordInvoker 批量重置弹性云服务器密码
func (c *EcsClient) BatchResetServersPasswordInvoker(request *model.BatchResetServersPasswordRequest) *BatchResetServersPasswordInvoker {
	requestDef := GenReqDefForBatchResetServersPassword()
	return &BatchResetServersPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartServers 批量启动云服务器
//
// 根据给定的云服务器ID列表，批量启动云服务器，一次最多可以启动1000台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchStartServers(request *model.BatchStartServersRequest) (*model.BatchStartServersResponse, error) {
	requestDef := GenReqDefForBatchStartServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartServersResponse), nil
	}
}

// BatchStartServersInvoker 批量启动云服务器
func (c *EcsClient) BatchStartServersInvoker(request *model.BatchStartServersRequest) *BatchStartServersInvoker {
	requestDef := GenReqDefForBatchStartServers()
	return &BatchStartServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopServers 批量关闭云服务器
//
// 根据给定的云服务器ID列表，批量关闭云服务器，一次最多可以关闭1000台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchStopServers(request *model.BatchStopServersRequest) (*model.BatchStopServersResponse, error) {
	requestDef := GenReqDefForBatchStopServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopServersResponse), nil
	}
}

// BatchStopServersInvoker 批量关闭云服务器
func (c *EcsClient) BatchStopServersInvoker(request *model.BatchStopServersRequest) *BatchStopServersInvoker {
	requestDef := GenReqDefForBatchStopServers()
	return &BatchStopServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateServersName 批量修改弹性云服务器
//
// 批量修改弹性云服务器信息。
// 当前仅支持批量修改云服务器名称，一次最多可以修改1000台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) BatchUpdateServersName(request *model.BatchUpdateServersNameRequest) (*model.BatchUpdateServersNameResponse, error) {
	requestDef := GenReqDefForBatchUpdateServersName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateServersNameResponse), nil
	}
}

// BatchUpdateServersNameInvoker 批量修改弹性云服务器
func (c *EcsClient) BatchUpdateServersNameInvoker(request *model.BatchUpdateServersNameRequest) *BatchUpdateServersNameInvoker {
	requestDef := GenReqDefForBatchUpdateServersName()
	return &BatchUpdateServersNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerChargeMode 更换云服务器计费模式
//
// 更换云服务器的计费模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ChangeServerChargeMode(request *model.ChangeServerChargeModeRequest) (*model.ChangeServerChargeModeResponse, error) {
	requestDef := GenReqDefForChangeServerChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerChargeModeResponse), nil
	}
}

// ChangeServerChargeModeInvoker 更换云服务器计费模式
func (c *EcsClient) ChangeServerChargeModeInvoker(request *model.ChangeServerChargeModeRequest) *ChangeServerChargeModeInvoker {
	requestDef := GenReqDefForChangeServerChargeMode()
	return &ChangeServerChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerNetworkInterface 更新云服务器指定网卡属性
//
// 更新云服务器指定网卡属性，当前仅支持更新网卡IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ChangeServerNetworkInterface(request *model.ChangeServerNetworkInterfaceRequest) (*model.ChangeServerNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForChangeServerNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerNetworkInterfaceResponse), nil
	}
}

// ChangeServerNetworkInterfaceInvoker 更新云服务器指定网卡属性
func (c *EcsClient) ChangeServerNetworkInterfaceInvoker(request *model.ChangeServerNetworkInterfaceRequest) *ChangeServerNetworkInterfaceInvoker {
	requestDef := GenReqDefForChangeServerNetworkInterface()
	return &ChangeServerNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerOsWithCloudInit 切换弹性云服务器操作系统(安装Cloud init)
//
// 切换弹性云服务器操作系统。支持弹性云服务器数据盘不变的情况下，使用新镜像重装系统盘。
//
// 调用该接口后，系统将卸载系统盘，然后使用新镜像重新创建系统盘，并挂载至弹性云服务器，实现切换操作系统功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ChangeServerOsWithCloudInit(request *model.ChangeServerOsWithCloudInitRequest) (*model.ChangeServerOsWithCloudInitResponse, error) {
	requestDef := GenReqDefForChangeServerOsWithCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerOsWithCloudInitResponse), nil
	}
}

// ChangeServerOsWithCloudInitInvoker 切换弹性云服务器操作系统(安装Cloud init)
func (c *EcsClient) ChangeServerOsWithCloudInitInvoker(request *model.ChangeServerOsWithCloudInitRequest) *ChangeServerOsWithCloudInitInvoker {
	requestDef := GenReqDefForChangeServerOsWithCloudInit()
	return &ChangeServerOsWithCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerOsWithoutCloudInit 切换弹性云服务器操作系统(未安装Cloud init)
//
// 切换弹性云服务器操作系统。
//
// 该接口支持未安装Cloud-init或Cloudbase-init的镜像使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ChangeServerOsWithoutCloudInit(request *model.ChangeServerOsWithoutCloudInitRequest) (*model.ChangeServerOsWithoutCloudInitResponse, error) {
	requestDef := GenReqDefForChangeServerOsWithoutCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerOsWithoutCloudInitResponse), nil
	}
}

// ChangeServerOsWithoutCloudInitInvoker 切换弹性云服务器操作系统(未安装Cloud init)
func (c *EcsClient) ChangeServerOsWithoutCloudInitInvoker(request *model.ChangeServerOsWithoutCloudInitRequest) *ChangeServerOsWithoutCloudInitInvoker {
	requestDef := GenReqDefForChangeServerOsWithoutCloudInit()
	return &ChangeServerOsWithoutCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeVpc 云服务器切换虚拟私有云
//
// 云服务器切换虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ChangeVpc(request *model.ChangeVpcRequest) (*model.ChangeVpcResponse, error) {
	requestDef := GenReqDefForChangeVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeVpcResponse), nil
	}
}

// ChangeVpcInvoker 云服务器切换虚拟私有云
func (c *EcsClient) ChangeVpcInvoker(request *model.ChangeVpcRequest) *ChangeVpcInvoker {
	requestDef := GenReqDefForChangeVpc()
	return &ChangeVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidServers 创建云服务器(按需)
//
// 创建一台或多台[按需付费](https://support.huaweicloud.com/productdesc-ecs/ecs_01_0065.html)方式的云服务器。
//
// 弹性云服务器的登录鉴权方式包括两种：密钥对、密码。为安全起见，推荐使用密钥对方式。
//
// - 密钥对
// 密钥对指使用密钥对作为弹性云服务器的鉴权方式。
// 接口调用方法：使用key_name字段，指定弹性云服务器登录时使用的密钥文件。
//
// - 密码
// 密码指使用设置初始密码方式作为弹性云服务器的鉴权方式，此时，您可以通过用户名密码方式登录弹性云服务器，Linux操作系统时为root用户的初始密码，Windows操作系统时为Administrator用户的初始密码。
//
// 接口调用方法：使用adminPass字段，指定管理员帐号的初始登录密码。对于镜像已安装Cloud-init的Linux云服务器，如果需要使用密文密码，可以使用user_data字段进行密码注入。
//
// &gt; 对于安装Cloud-init镜像的Linux云服务器云主机，若指定user_data字段，则adminPass字段无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) CreatePostPaidServers(request *model.CreatePostPaidServersRequest) (*model.CreatePostPaidServersResponse, error) {
	requestDef := GenReqDefForCreatePostPaidServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidServersResponse), nil
	}
}

// CreatePostPaidServersInvoker 创建云服务器(按需)
func (c *EcsClient) CreatePostPaidServersInvoker(request *model.CreatePostPaidServersRequest) *CreatePostPaidServersInvoker {
	requestDef := GenReqDefForCreatePostPaidServers()
	return &CreatePostPaidServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServerGroup 创建云服务器组
//
// 创建弹性云服务器组。
//
// 与原生的创建云服务器组接口不同之处在于该接口支持企业项目细粒度权限的校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) CreateServerGroup(request *model.CreateServerGroupRequest) (*model.CreateServerGroupResponse, error) {
	requestDef := GenReqDefForCreateServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServerGroupResponse), nil
	}
}

// CreateServerGroupInvoker 创建云服务器组
func (c *EcsClient) CreateServerGroupInvoker(request *model.CreateServerGroupRequest) *CreateServerGroupInvoker {
	requestDef := GenReqDefForCreateServerGroup()
	return &CreateServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServers 创建云服务器
//
// 创建一台或多台云服务器。
//
// 指该接口兼容《弹性云服务器接口参考》创建云服务器v1的功能，同时合入新功能，支持创建[包年/包月](https://support.huaweicloud.com/productdesc-ecs/ecs_01_0065.html)的弹性云服务器。
//
// 弹性云服务器的登录鉴权方式包括两种：密钥对、密码。为安全起见，推荐使用密钥对方式。
//
// - 密钥对
//
// 指使用密钥对作为弹性云服务器的鉴权方式。
//
// 接口调用方法：使用key_name字段，指定弹性云服务器登录时使用的密钥文件。
//
// - 密码
//
// 指使用设置初始密码方式作为弹性云服务器的鉴权方式，此时，您可以通过用户名密码方式登录弹性云服务器，Linux操作系统时为root用户的初始密码，Windows操作系统时为Administrator用户的初始密码。
//
// 接口调用方法：使用adminPass字段，指定管理员帐号的初始登录密码。对于镜像已安装Cloud-init的Linux云服务器，如果需要使用密文密码，可以使用user_data字段进行密码注入。
//
// &gt; 对于安装Cloud-init镜像的Linux云服务器云主机，若指定user_data字段，则adminPass字段无效。
//
// 购买操作示例：
// - [使用API购买ECS过程中常见问题及处理方法](https://support.huaweicloud.com/api-ecs/ecs_04_0007.html)
// - [获取Token并检验Token的有效期 ](https://support.huaweicloud.com/api-ecs/ecs_04_0008.html)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) CreateServers(request *model.CreateServersRequest) (*model.CreateServersResponse, error) {
	requestDef := GenReqDefForCreateServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServersResponse), nil
	}
}

// CreateServersInvoker 创建云服务器
func (c *EcsClient) CreateServersInvoker(request *model.CreateServersRequest) *CreateServersInvoker {
	requestDef := GenReqDefForCreateServers()
	return &CreateServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerGroup 删除云服务器组
//
// 删除云服务器组。
//
// 与原生的删除云服务器组接口不同之处在于该接口支持企业项目细粒度权限的校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DeleteServerGroup(request *model.DeleteServerGroupRequest) (*model.DeleteServerGroupResponse, error) {
	requestDef := GenReqDefForDeleteServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupResponse), nil
	}
}

// DeleteServerGroupInvoker 删除云服务器组
func (c *EcsClient) DeleteServerGroupInvoker(request *model.DeleteServerGroupRequest) *DeleteServerGroupInvoker {
	requestDef := GenReqDefForDeleteServerGroup()
	return &DeleteServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerGroupMember 删除云服务器组成员
//
// 将弹性云服务器移出云服务器组。移出后，该云服务器与云服务器组中的成员不再遵从反亲和策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DeleteServerGroupMember(request *model.DeleteServerGroupMemberRequest) (*model.DeleteServerGroupMemberResponse, error) {
	requestDef := GenReqDefForDeleteServerGroupMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupMemberResponse), nil
	}
}

// DeleteServerGroupMemberInvoker 删除云服务器组成员
func (c *EcsClient) DeleteServerGroupMemberInvoker(request *model.DeleteServerGroupMemberRequest) *DeleteServerGroupMemberInvoker {
	requestDef := GenReqDefForDeleteServerGroupMember()
	return &DeleteServerGroupMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerMetadata 删除云服务器指定元数据
//
// 删除云服务器指定元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DeleteServerMetadata(request *model.DeleteServerMetadataRequest) (*model.DeleteServerMetadataResponse, error) {
	requestDef := GenReqDefForDeleteServerMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerMetadataResponse), nil
	}
}

// DeleteServerMetadataInvoker 删除云服务器指定元数据
func (c *EcsClient) DeleteServerMetadataInvoker(request *model.DeleteServerMetadataRequest) *DeleteServerMetadataInvoker {
	requestDef := GenReqDefForDeleteServerMetadata()
	return &DeleteServerMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerPassword 云服务器清除密码(企业项目)
//
// 清除Windows云服务器初始安装时系统生成的密码记录。清除密码后，不影响云服务器密码登录功能，但不能再使用获取密码功能来查询该云服务器密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DeleteServerPassword(request *model.DeleteServerPasswordRequest) (*model.DeleteServerPasswordResponse, error) {
	requestDef := GenReqDefForDeleteServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerPasswordResponse), nil
	}
}

// DeleteServerPasswordInvoker 云服务器清除密码(企业项目)
func (c *EcsClient) DeleteServerPasswordInvoker(request *model.DeleteServerPasswordRequest) *DeleteServerPasswordInvoker {
	requestDef := GenReqDefForDeleteServerPassword()
	return &DeleteServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServers 删除云服务器
//
// 根据指定的云服务器ID列表，删除云服务器。
//
// 系统支持删除单台云服务器和批量删除多台云服务器操作，批量删除云服务器时，一次最多可以删除1000台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DeleteServers(request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	requestDef := GenReqDefForDeleteServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServersResponse), nil
	}
}

// DeleteServersInvoker 删除云服务器
func (c *EcsClient) DeleteServersInvoker(request *model.DeleteServersRequest) *DeleteServersInvoker {
	requestDef := GenReqDefForDeleteServers()
	return &DeleteServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachServerVolume 弹性云服务器卸载磁盘
//
// 从弹性云服务器中卸载磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DetachServerVolume(request *model.DetachServerVolumeRequest) (*model.DetachServerVolumeResponse, error) {
	requestDef := GenReqDefForDetachServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachServerVolumeResponse), nil
	}
}

// DetachServerVolumeInvoker 弹性云服务器卸载磁盘
func (c *EcsClient) DetachServerVolumeInvoker(request *model.DetachServerVolumeRequest) *DetachServerVolumeInvoker {
	requestDef := GenReqDefForDetachServerVolume()
	return &DetachServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateServerVirtualIp 云服务器网卡解绑虚拟IP地址
//
// 虚拟IP地址用于为网卡提供第二个IP地址，同时支持与多个弹性云服务器的网卡绑定，从而实现多个弹性云服务器之间的高可用性。
//
// 该接口用于解绑定弹性云服务器网卡的虚拟IP地址。解绑后，网卡不会被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) DisassociateServerVirtualIp(request *model.DisassociateServerVirtualIpRequest) (*model.DisassociateServerVirtualIpResponse, error) {
	requestDef := GenReqDefForDisassociateServerVirtualIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateServerVirtualIpResponse), nil
	}
}

// DisassociateServerVirtualIpInvoker 云服务器网卡解绑虚拟IP地址
func (c *EcsClient) DisassociateServerVirtualIpInvoker(request *model.DisassociateServerVirtualIpRequest) *DisassociateServerVirtualIpInvoker {
	requestDef := GenReqDefForDisassociateServerVirtualIp()
	return &DisassociateServerVirtualIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudServers 查询云服务器列表接口
//
// 查询云服务器列表接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListCloudServers(request *model.ListCloudServersRequest) (*model.ListCloudServersResponse, error) {
	requestDef := GenReqDefForListCloudServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudServersResponse), nil
	}
}

// ListCloudServersInvoker 查询云服务器列表接口
func (c *EcsClient) ListCloudServersInvoker(request *model.ListCloudServersRequest) *ListCloudServersInvoker {
	requestDef := GenReqDefForListCloudServers()
	return &ListCloudServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavorSellPolicies 查询规格销售策略
//
// 查询规格销售策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListFlavorSellPolicies(request *model.ListFlavorSellPoliciesRequest) (*model.ListFlavorSellPoliciesResponse, error) {
	requestDef := GenReqDefForListFlavorSellPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorSellPoliciesResponse), nil
	}
}

// ListFlavorSellPoliciesInvoker 查询规格销售策略
func (c *EcsClient) ListFlavorSellPoliciesInvoker(request *model.ListFlavorSellPoliciesRequest) *ListFlavorSellPoliciesInvoker {
	requestDef := GenReqDefForListFlavorSellPolicies()
	return &ListFlavorSellPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询规格详情和规格扩展信息列表
//
// 查询云服务器规格详情信息和规格扩展信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询规格详情和规格扩展信息列表
func (c *EcsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResizeFlavors 查询云服务器规格变更支持列表
//
// 变更规格时，部分规格的云服务器之间不能互相变更。您可以通过本接口，通过指定弹性云服务器规格，查询该规格可以变更的规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListResizeFlavors(request *model.ListResizeFlavorsRequest) (*model.ListResizeFlavorsResponse, error) {
	requestDef := GenReqDefForListResizeFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResizeFlavorsResponse), nil
	}
}

// ListResizeFlavorsInvoker 查询云服务器规格变更支持列表
func (c *EcsClient) ListResizeFlavorsInvoker(request *model.ListResizeFlavorsRequest) *ListResizeFlavorsInvoker {
	requestDef := GenReqDefForListResizeFlavors()
	return &ListResizeFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerBlockDevices 查询弹性云服务器挂载磁盘列表详情信息
//
// 查询弹性云服务器挂载的磁盘信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServerBlockDevices(request *model.ListServerBlockDevicesRequest) (*model.ListServerBlockDevicesResponse, error) {
	requestDef := GenReqDefForListServerBlockDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerBlockDevicesResponse), nil
	}
}

// ListServerBlockDevicesInvoker 查询弹性云服务器挂载磁盘列表详情信息
func (c *EcsClient) ListServerBlockDevicesInvoker(request *model.ListServerBlockDevicesRequest) *ListServerBlockDevicesInvoker {
	requestDef := GenReqDefForListServerBlockDevices()
	return &ListServerBlockDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerGroups 查询云服务器组列表
//
// 查询弹性云服务器组。
//
// 与原生的创建云服务器组接口不同之处在于该接口支持企业项目细粒度权限的校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServerGroups(request *model.ListServerGroupsRequest) (*model.ListServerGroupsResponse, error) {
	requestDef := GenReqDefForListServerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerGroupsResponse), nil
	}
}

// ListServerGroupsInvoker 查询云服务器组列表
func (c *EcsClient) ListServerGroupsInvoker(request *model.ListServerGroupsRequest) *ListServerGroupsInvoker {
	requestDef := GenReqDefForListServerGroups()
	return &ListServerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerInterfaces 查询云服务器网卡信息
//
// 查询云服务器网卡信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServerInterfaces(request *model.ListServerInterfacesRequest) (*model.ListServerInterfacesResponse, error) {
	requestDef := GenReqDefForListServerInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerInterfacesResponse), nil
	}
}

// ListServerInterfacesInvoker 查询云服务器网卡信息
func (c *EcsClient) ListServerInterfacesInvoker(request *model.ListServerInterfacesRequest) *ListServerInterfacesInvoker {
	requestDef := GenReqDefForListServerInterfaces()
	return &ListServerInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerTags 查询项目标签
//
// 项目（Project）用于将OpenStack的资源（计算资源、存储资源和网络资源）进行分组和隔离。项目可以是一个部门或者一个项目组。一个帐户中可以创建多个项目。
//
// 该接口用于查询用户在指定项目所使用的全部标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServerTags(request *model.ListServerTagsRequest) (*model.ListServerTagsResponse, error) {
	requestDef := GenReqDefForListServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerTagsResponse), nil
	}
}

// ListServerTagsInvoker 查询项目标签
func (c *EcsClient) ListServerTagsInvoker(request *model.ListServerTagsRequest) *ListServerTagsInvoker {
	requestDef := GenReqDefForListServerTags()
	return &ListServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListServersByTag 按标签查询云服务器列表
//
// 使用标签过滤弹性云服务器，并返回云服务器使用的所有标签和资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServersByTag(request *model.ListServersByTagRequest) (*model.ListServersByTagResponse, error) {
	requestDef := GenReqDefForListServersByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersByTagResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListServersByTagInvoker 按标签查询云服务器列表
func (c *EcsClient) ListServersByTagInvoker(request *model.ListServersByTagRequest) *ListServersByTagInvoker {
	requestDef := GenReqDefForListServersByTag()
	return &ListServersByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServersDetails 查询云服务器详情列表
//
// 根据用户请求条件从数据库筛选、查询所有的弹性云服务器，并关联相关表获取到弹性云服务器的详细信息。
//
// 该接口支持查询弹性云服务器计费方式，以及是否被冻结。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ListServersDetails(request *model.ListServersDetailsRequest) (*model.ListServersDetailsResponse, error) {
	requestDef := GenReqDefForListServersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersDetailsResponse), nil
	}
}

// ListServersDetailsInvoker 查询云服务器详情列表
func (c *EcsClient) ListServersDetailsInvoker(request *model.ListServersDetailsRequest) *ListServersDetailsInvoker {
	requestDef := GenReqDefForListServersDetails()
	return &ListServersDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateServer 冷迁移云服务器
//
// - 将部署在专属主机上的弹性云服务器迁移至其他专属主机。
// - 将部署在专属主机上的弹性云服务器迁移至公共资源池，即不再部署在专属主机上。
// - 将公共资源池的弹性云服务器迁移至专属主机上，成为专属主机上部署的弹性云服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) MigrateServer(request *model.MigrateServerRequest) (*model.MigrateServerResponse, error) {
	requestDef := GenReqDefForMigrateServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateServerResponse), nil
	}
}

// MigrateServerInvoker 冷迁移云服务器
func (c *EcsClient) MigrateServerInvoker(request *model.MigrateServerRequest) *MigrateServerInvoker {
	requestDef := GenReqDefForMigrateServer()
	return &MigrateServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaAssociateSecurityGroup 添加安全组
//
// 为弹性云服务器添加一个安全组。
//
// 添加多个安全组时，建议最多为弹性云服务器添加5个安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaAssociateSecurityGroup(request *model.NovaAssociateSecurityGroupRequest) (*model.NovaAssociateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNovaAssociateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaAssociateSecurityGroupResponse), nil
	}
}

// NovaAssociateSecurityGroupInvoker 添加安全组
func (c *EcsClient) NovaAssociateSecurityGroupInvoker(request *model.NovaAssociateSecurityGroupRequest) *NovaAssociateSecurityGroupInvoker {
	requestDef := GenReqDefForNovaAssociateSecurityGroup()
	return &NovaAssociateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaAttachInterface 添加云服务器网卡
//
// 给云服务器添加一张网卡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaAttachInterface(request *model.NovaAttachInterfaceRequest) (*model.NovaAttachInterfaceResponse, error) {
	requestDef := GenReqDefForNovaAttachInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaAttachInterfaceResponse), nil
	}
}

// NovaAttachInterfaceInvoker 添加云服务器网卡
func (c *EcsClient) NovaAttachInterfaceInvoker(request *model.NovaAttachInterfaceRequest) *NovaAttachInterfaceInvoker {
	requestDef := GenReqDefForNovaAttachInterface()
	return &NovaAttachInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaCreateKeypair 创建和导入SSH密钥
//
// 创建SSH密钥，或把公钥导入系统，生成密钥对。
//
// 创建SSH密钥成功后，请把响应数据中的私钥内容保存到本地文件，用户使用该私钥登录云服务器云主机。为保证云服务器云主机器安全，私钥数据只能读取一次，请妥善保管。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaCreateKeypair(request *model.NovaCreateKeypairRequest) (*model.NovaCreateKeypairResponse, error) {
	requestDef := GenReqDefForNovaCreateKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaCreateKeypairResponse), nil
	}
}

// NovaCreateKeypairInvoker 创建和导入SSH密钥
func (c *EcsClient) NovaCreateKeypairInvoker(request *model.NovaCreateKeypairRequest) *NovaCreateKeypairInvoker {
	requestDef := GenReqDefForNovaCreateKeypair()
	return &NovaCreateKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaCreateServers 创建云服务器
//
// 创建一台弹性云服务器。
//
// 弹性云服务器创建完成后，如需开启自动恢复功能，可以调用配置云服务器自动恢复的接口，具体使用请参见管理云服务器自动恢复动作。
//
// 该接口在云服务器创建失败后不支持自动回滚。若需要自动回滚能力，可以调用POST /v1/{project_id}/cloudservers接口，具体使用请参见创建云服务器（按需）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaCreateServers(request *model.NovaCreateServersRequest) (*model.NovaCreateServersResponse, error) {
	requestDef := GenReqDefForNovaCreateServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaCreateServersResponse), nil
	}
}

// NovaCreateServersInvoker 创建云服务器
func (c *EcsClient) NovaCreateServersInvoker(request *model.NovaCreateServersRequest) *NovaCreateServersInvoker {
	requestDef := GenReqDefForNovaCreateServers()
	return &NovaCreateServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaDeleteKeypair 删除SSH密钥
//
// 根据SSH密钥的名称，删除指定SSH密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaDeleteKeypair(request *model.NovaDeleteKeypairRequest) (*model.NovaDeleteKeypairResponse, error) {
	requestDef := GenReqDefForNovaDeleteKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDeleteKeypairResponse), nil
	}
}

// NovaDeleteKeypairInvoker 删除SSH密钥
func (c *EcsClient) NovaDeleteKeypairInvoker(request *model.NovaDeleteKeypairRequest) *NovaDeleteKeypairInvoker {
	requestDef := GenReqDefForNovaDeleteKeypair()
	return &NovaDeleteKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaDeleteServer 删除云服务器
//
// 删除一台云服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaDeleteServer(request *model.NovaDeleteServerRequest) (*model.NovaDeleteServerResponse, error) {
	requestDef := GenReqDefForNovaDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDeleteServerResponse), nil
	}
}

// NovaDeleteServerInvoker 删除云服务器
func (c *EcsClient) NovaDeleteServerInvoker(request *model.NovaDeleteServerRequest) *NovaDeleteServerInvoker {
	requestDef := GenReqDefForNovaDeleteServer()
	return &NovaDeleteServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaDisassociateSecurityGroup 移除安全组
//
// 移除弹性云服务器中的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaDisassociateSecurityGroup(request *model.NovaDisassociateSecurityGroupRequest) (*model.NovaDisassociateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNovaDisassociateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaDisassociateSecurityGroupResponse), nil
	}
}

// NovaDisassociateSecurityGroupInvoker 移除安全组
func (c *EcsClient) NovaDisassociateSecurityGroupInvoker(request *model.NovaDisassociateSecurityGroupRequest) *NovaDisassociateSecurityGroupInvoker {
	requestDef := GenReqDefForNovaDisassociateSecurityGroup()
	return &NovaDisassociateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaListAvailabilityZones 查询可用区列表
//
// 查询可用域列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaListAvailabilityZones(request *model.NovaListAvailabilityZonesRequest) (*model.NovaListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForNovaListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListAvailabilityZonesResponse), nil
	}
}

// NovaListAvailabilityZonesInvoker 查询可用区列表
func (c *EcsClient) NovaListAvailabilityZonesInvoker(request *model.NovaListAvailabilityZonesRequest) *NovaListAvailabilityZonesInvoker {
	requestDef := GenReqDefForNovaListAvailabilityZones()
	return &NovaListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaListKeypairs 查询SSH密钥列表
//
// 查询SSH密钥信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaListKeypairs(request *model.NovaListKeypairsRequest) (*model.NovaListKeypairsResponse, error) {
	requestDef := GenReqDefForNovaListKeypairs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListKeypairsResponse), nil
	}
}

// NovaListKeypairsInvoker 查询SSH密钥列表
func (c *EcsClient) NovaListKeypairsInvoker(request *model.NovaListKeypairsRequest) *NovaListKeypairsInvoker {
	requestDef := GenReqDefForNovaListKeypairs()
	return &NovaListKeypairsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaListServerSecurityGroups 查询指定云服务器安全组列表
//
// 查询指定弹性云服务器的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaListServerSecurityGroups(request *model.NovaListServerSecurityGroupsRequest) (*model.NovaListServerSecurityGroupsResponse, error) {
	requestDef := GenReqDefForNovaListServerSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListServerSecurityGroupsResponse), nil
	}
}

// NovaListServerSecurityGroupsInvoker 查询指定云服务器安全组列表
func (c *EcsClient) NovaListServerSecurityGroupsInvoker(request *model.NovaListServerSecurityGroupsRequest) *NovaListServerSecurityGroupsInvoker {
	requestDef := GenReqDefForNovaListServerSecurityGroups()
	return &NovaListServerSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaListServersDetails 查询云服务器详情列表
//
// 查询云服务器详情信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaListServersDetails(request *model.NovaListServersDetailsRequest) (*model.NovaListServersDetailsResponse, error) {
	requestDef := GenReqDefForNovaListServersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListServersDetailsResponse), nil
	}
}

// NovaListServersDetailsInvoker 查询云服务器详情列表
func (c *EcsClient) NovaListServersDetailsInvoker(request *model.NovaListServersDetailsRequest) *NovaListServersDetailsInvoker {
	requestDef := GenReqDefForNovaListServersDetails()
	return &NovaListServersDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaShowKeypair 查询SSH密钥详情
//
// 根据SSH密钥名称查询指定SSH密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaShowKeypair(request *model.NovaShowKeypairRequest) (*model.NovaShowKeypairResponse, error) {
	requestDef := GenReqDefForNovaShowKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowKeypairResponse), nil
	}
}

// NovaShowKeypairInvoker 查询SSH密钥详情
func (c *EcsClient) NovaShowKeypairInvoker(request *model.NovaShowKeypairRequest) *NovaShowKeypairInvoker {
	requestDef := GenReqDefForNovaShowKeypair()
	return &NovaShowKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaShowServer 查询云服务器详情
//
// 根据云服务器ID，查询云服务器的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaShowServer(request *model.NovaShowServerRequest) (*model.NovaShowServerResponse, error) {
	requestDef := GenReqDefForNovaShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowServerResponse), nil
	}
}

// NovaShowServerInvoker 查询云服务器详情
func (c *EcsClient) NovaShowServerInvoker(request *model.NovaShowServerRequest) *NovaShowServerInvoker {
	requestDef := GenReqDefForNovaShowServer()
	return &NovaShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaShowServerInterface 查询指定云服务器网卡信息
//
// 根据网卡ID，查询云服务器网卡信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaShowServerInterface(request *model.NovaShowServerInterfaceRequest) (*model.NovaShowServerInterfaceResponse, error) {
	requestDef := GenReqDefForNovaShowServerInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowServerInterfaceResponse), nil
	}
}

// NovaShowServerInterfaceInvoker 查询指定云服务器网卡信息
func (c *EcsClient) NovaShowServerInterfaceInvoker(request *model.NovaShowServerInterfaceRequest) *NovaShowServerInterfaceInvoker {
	requestDef := GenReqDefForNovaShowServerInterface()
	return &NovaShowServerInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterServerMonitor 注册云服务器监控
//
// 将云服务器添加到监控表中。
//
// 注册到监控表中的云服务会被ceilometer周期性采集监控数据，包括平台的版本、cpu信息、内存、网卡、磁盘、硬件平台等信息，这些数据上报给云监控。例如SAP云服务器内部的插件会周期性从云监控中查询监控数据，以报表形式呈现给SAP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) RegisterServerMonitor(request *model.RegisterServerMonitorRequest) (*model.RegisterServerMonitorResponse, error) {
	requestDef := GenReqDefForRegisterServerMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterServerMonitorResponse), nil
	}
}

// RegisterServerMonitorInvoker 注册云服务器监控
func (c *EcsClient) RegisterServerMonitorInvoker(request *model.RegisterServerMonitorRequest) *RegisterServerMonitorInvoker {
	requestDef := GenReqDefForRegisterServerMonitor()
	return &RegisterServerMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallServerWithCloudInit 重装弹性云服务器操作系统(安装Cloud-init)
//
// 重装弹性云服务器的操作系统。支持弹性云服务器数据盘不变的情况下，使用原镜像重装系统盘。
//
// 调用该接口后，系统将卸载系统盘，然后使用原镜像重新创建系统盘，并挂载至弹性云服务器，实现重装操作系统功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ReinstallServerWithCloudInit(request *model.ReinstallServerWithCloudInitRequest) (*model.ReinstallServerWithCloudInitResponse, error) {
	requestDef := GenReqDefForReinstallServerWithCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallServerWithCloudInitResponse), nil
	}
}

// ReinstallServerWithCloudInitInvoker 重装弹性云服务器操作系统(安装Cloud-init)
func (c *EcsClient) ReinstallServerWithCloudInitInvoker(request *model.ReinstallServerWithCloudInitRequest) *ReinstallServerWithCloudInitInvoker {
	requestDef := GenReqDefForReinstallServerWithCloudInit()
	return &ReinstallServerWithCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallServerWithoutCloudInit 重装弹性云服务器操作系统(未安装Cloud init)
//
// 重装弹性云服务器的操作系统。
//
// 该接口支持未安装Cloud-init或Cloudbase-init的镜像。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ReinstallServerWithoutCloudInit(request *model.ReinstallServerWithoutCloudInitRequest) (*model.ReinstallServerWithoutCloudInitResponse, error) {
	requestDef := GenReqDefForReinstallServerWithoutCloudInit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallServerWithoutCloudInitResponse), nil
	}
}

// ReinstallServerWithoutCloudInitInvoker 重装弹性云服务器操作系统(未安装Cloud init)
func (c *EcsClient) ReinstallServerWithoutCloudInitInvoker(request *model.ReinstallServerWithoutCloudInitRequest) *ReinstallServerWithoutCloudInitInvoker {
	requestDef := GenReqDefForReinstallServerWithoutCloudInit()
	return &ReinstallServerWithoutCloudInitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetServerPassword 一键重置弹性云服务器密码(企业项目)
//
// 重置弹性云服务器管理帐号（root用户或Administrator用户）的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ResetServerPassword(request *model.ResetServerPasswordRequest) (*model.ResetServerPasswordResponse, error) {
	requestDef := GenReqDefForResetServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetServerPasswordResponse), nil
	}
}

// ResetServerPasswordInvoker 一键重置弹性云服务器密码(企业项目)
func (c *EcsClient) ResetServerPasswordInvoker(request *model.ResetServerPasswordRequest) *ResetServerPasswordInvoker {
	requestDef := GenReqDefForResetServerPassword()
	return &ResetServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizePostPaidServer 变更云服务器规格(按需)
//
// 当您创建的弹性云服务器规格无法满足业务需要时，可以变更云服务器规格，升级vCPU、内存。具体接口的使用，请参见本节内容。
//
// 变更规格时，部分规格的云服务器之间不能互相变更。
//
// 您可以通过接口“/v1/{project_id}/cloudservers/resize_flavors?{instance_uuid,source_flavor_id,source_flavor_name}”查询支持列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ResizePostPaidServer(request *model.ResizePostPaidServerRequest) (*model.ResizePostPaidServerResponse, error) {
	requestDef := GenReqDefForResizePostPaidServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizePostPaidServerResponse), nil
	}
}

// ResizePostPaidServerInvoker 变更云服务器规格(按需)
func (c *EcsClient) ResizePostPaidServerInvoker(request *model.ResizePostPaidServerRequest) *ResizePostPaidServerInvoker {
	requestDef := GenReqDefForResizePostPaidServer()
	return &ResizePostPaidServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeServer 变更云服务器规格
//
// 变更云服务器规格。
//
// v1.1版本：指该接口兼容v1接口的功能，同时合入新功能，支持变更包年/包月弹性云服务器的规格。
//
// 注意事项：
//
// - 该接口可以使用合作伙伴自身的AK/SK或者token调用，也可以用合作伙伴子客户的AK/SK或者token来调用。
// - 如果使用AK/SK认证方式，示例代码中region请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)中“弹性云服务 ECS”下“区域”的内容，，serviceName（英文服务名称缩写）请指定为ECS。
// - Endpoint请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)中“弹性云服务 ECS”下“终端节点（Endpoint）”的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ResizeServer(request *model.ResizeServerRequest) (*model.ResizeServerResponse, error) {
	requestDef := GenReqDefForResizeServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeServerResponse), nil
	}
}

// ResizeServerInvoker 变更云服务器规格
func (c *EcsClient) ResizeServerInvoker(request *model.ResizeServerRequest) *ResizeServerInvoker {
	requestDef := GenReqDefForResizeServer()
	return &ResizeServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResetPasswordFlag 查询是否支持一键重置密码
//
// 查询弹性云服务器是否支持一键重置密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowResetPasswordFlag(request *model.ShowResetPasswordFlagRequest) (*model.ShowResetPasswordFlagResponse, error) {
	requestDef := GenReqDefForShowResetPasswordFlag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResetPasswordFlagResponse), nil
	}
}

// ShowResetPasswordFlagInvoker 查询是否支持一键重置密码
func (c *EcsClient) ShowResetPasswordFlagInvoker(request *model.ShowResetPasswordFlagRequest) *ShowResetPasswordFlagInvoker {
	requestDef := GenReqDefForShowResetPasswordFlag()
	return &ShowResetPasswordFlagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServer 查询云服务器详情
//
// 查询弹性云服务器的详细信息。
//
// 该接口支持查询弹性云服务器的计费方式，以及是否被冻结。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

// ShowServerInvoker 查询云服务器详情
func (c *EcsClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerBlockDevice 查询弹性云服务器单个磁盘信息
//
// 查询弹性云服务器挂载的单个磁盘信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerBlockDevice(request *model.ShowServerBlockDeviceRequest) (*model.ShowServerBlockDeviceResponse, error) {
	requestDef := GenReqDefForShowServerBlockDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerBlockDeviceResponse), nil
	}
}

// ShowServerBlockDeviceInvoker 查询弹性云服务器单个磁盘信息
func (c *EcsClient) ShowServerBlockDeviceInvoker(request *model.ShowServerBlockDeviceRequest) *ShowServerBlockDeviceInvoker {
	requestDef := GenReqDefForShowServerBlockDevice()
	return &ShowServerBlockDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerGroup 查询云服务器组详情
//
// 查询弹性云服务器组详情。
//
// 与原生的创建云服务器组接口不同之处在于该接口支持企业项目细粒度权限的校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerGroup(request *model.ShowServerGroupRequest) (*model.ShowServerGroupResponse, error) {
	requestDef := GenReqDefForShowServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupResponse), nil
	}
}

// ShowServerGroupInvoker 查询云服务器组详情
func (c *EcsClient) ShowServerGroupInvoker(request *model.ShowServerGroupRequest) *ShowServerGroupInvoker {
	requestDef := GenReqDefForShowServerGroup()
	return &ShowServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerLimits 查询租户配额
//
// 查询租户配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerLimits(request *model.ShowServerLimitsRequest) (*model.ShowServerLimitsResponse, error) {
	requestDef := GenReqDefForShowServerLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerLimitsResponse), nil
	}
}

// ShowServerLimitsInvoker 查询租户配额
func (c *EcsClient) ShowServerLimitsInvoker(request *model.ShowServerLimitsRequest) *ShowServerLimitsInvoker {
	requestDef := GenReqDefForShowServerLimits()
	return &ShowServerLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerPassword 云服务器获取密码(企业项目)
//
// 当通过支持Cloudbase-init功能的镜像创建Windows云服务器时，获取云服务器初始安装时系统生成的管理员帐户（Administrator帐户或Cloudbase-init设置的帐户）随机密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerPassword(request *model.ShowServerPasswordRequest) (*model.ShowServerPasswordResponse, error) {
	requestDef := GenReqDefForShowServerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerPasswordResponse), nil
	}
}

// ShowServerPasswordInvoker 云服务器获取密码(企业项目)
func (c *EcsClient) ShowServerPasswordInvoker(request *model.ShowServerPasswordRequest) *ShowServerPasswordInvoker {
	requestDef := GenReqDefForShowServerPassword()
	return &ShowServerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerRemoteConsole 获取VNC远程登录地址
//
// 获取弹性云服务器VNC远程登录地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerRemoteConsole(request *model.ShowServerRemoteConsoleRequest) (*model.ShowServerRemoteConsoleResponse, error) {
	requestDef := GenReqDefForShowServerRemoteConsole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerRemoteConsoleResponse), nil
	}
}

// ShowServerRemoteConsoleInvoker 获取VNC远程登录地址
func (c *EcsClient) ShowServerRemoteConsoleInvoker(request *model.ShowServerRemoteConsoleRequest) *ShowServerRemoteConsoleInvoker {
	requestDef := GenReqDefForShowServerRemoteConsole()
	return &ShowServerRemoteConsoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerTags 查询云服务器标签
//
// - 查询指定云服务器的标签信息。
//
// - 标签管理服务TMS使用该接口查询指定云服务器的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowServerTags(request *model.ShowServerTagsRequest) (*model.ShowServerTagsResponse, error) {
	requestDef := GenReqDefForShowServerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerTagsResponse), nil
	}
}

// ShowServerTagsInvoker 查询云服务器标签
func (c *EcsClient) ShowServerTagsInvoker(request *model.ShowServerTagsRequest) *ShowServerTagsInvoker {
	requestDef := GenReqDefForShowServerTags()
	return &ShowServerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServer 修改云服务器
//
// 修改云服务器信息，目前支持修改云服务器名称及描述和hostname。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) UpdateServer(request *model.UpdateServerRequest) (*model.UpdateServerResponse, error) {
	requestDef := GenReqDefForUpdateServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerResponse), nil
	}
}

// UpdateServerInvoker 修改云服务器
func (c *EcsClient) UpdateServerInvoker(request *model.UpdateServerRequest) *UpdateServerInvoker {
	requestDef := GenReqDefForUpdateServer()
	return &UpdateServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerAutoTerminateTime 修改云服务器定时删除时间
//
// 修改按需服务器，设置定时删除时间。如果设置的定时删除时间为空字符串，表示取消定时删除。
//
// 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ecs:cloudServers:put。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) UpdateServerAutoTerminateTime(request *model.UpdateServerAutoTerminateTimeRequest) (*model.UpdateServerAutoTerminateTimeResponse, error) {
	requestDef := GenReqDefForUpdateServerAutoTerminateTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerAutoTerminateTimeResponse), nil
	}
}

// UpdateServerAutoTerminateTimeInvoker 修改云服务器定时删除时间
func (c *EcsClient) UpdateServerAutoTerminateTimeInvoker(request *model.UpdateServerAutoTerminateTimeRequest) *UpdateServerAutoTerminateTimeInvoker {
	requestDef := GenReqDefForUpdateServerAutoTerminateTime()
	return &UpdateServerAutoTerminateTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerBlockDevice 修改云服务器挂载的单个磁盘信息
//
// 修改云服务器云主机挂载的单个磁盘信息。&#39;当前仅支持修改delete_on_termination字段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) UpdateServerBlockDevice(request *model.UpdateServerBlockDeviceRequest) (*model.UpdateServerBlockDeviceResponse, error) {
	requestDef := GenReqDefForUpdateServerBlockDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerBlockDeviceResponse), nil
	}
}

// UpdateServerBlockDeviceInvoker 修改云服务器挂载的单个磁盘信息
func (c *EcsClient) UpdateServerBlockDeviceInvoker(request *model.UpdateServerBlockDeviceRequest) *UpdateServerBlockDeviceInvoker {
	requestDef := GenReqDefForUpdateServerBlockDevice()
	return &UpdateServerBlockDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerMetadata 更新云服务器元数据
//
// 更新云服务器元数据。
//
// - 如果元数据中没有待更新字段，则自动添加该字段。
//
// - 如果元数据中已存在待更新字段，则直接更新字段值。
//
// - 如果元数据中的字段不再请求参数中，则保持不变
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) UpdateServerMetadata(request *model.UpdateServerMetadataRequest) (*model.UpdateServerMetadataResponse, error) {
	requestDef := GenReqDefForUpdateServerMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerMetadataResponse), nil
	}
}

// UpdateServerMetadataInvoker 更新云服务器元数据
func (c *EcsClient) UpdateServerMetadataInvoker(request *model.UpdateServerMetadataRequest) *UpdateServerMetadataInvoker {
	requestDef := GenReqDefForUpdateServerMetadata()
	return &UpdateServerMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaListVersions 查询API版本信息列表
//
// 返回Nova当前所有可用的版本。
//
// 为了支持功能不断扩展，Nova API支持版本号区分。Nova中有两种形式的版本号：
//
// - \&quot;主版本号\&quot;: 具有独立的url。
// - \&quot;微版本号\&quot;: 通过Http请求头X-OpenStack-Nova-API-Version来使用，从2.27版本后更改为OpenStack-API-Version。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaListVersions(request *model.NovaListVersionsRequest) (*model.NovaListVersionsResponse, error) {
	requestDef := GenReqDefForNovaListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaListVersionsResponse), nil
	}
}

// NovaListVersionsInvoker 查询API版本信息列表
func (c *EcsClient) NovaListVersionsInvoker(request *model.NovaListVersionsRequest) *NovaListVersionsInvoker {
	requestDef := GenReqDefForNovaListVersions()
	return &NovaListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NovaShowVersion 查询指定API版本信息
//
// 返回指定版本的信息。
// 为了支持功能不断扩展，Nova API支持版本号区分。Nova中有两种形式的版本号：
//
// - \&quot;主版本号\&quot;: 具有独立的url。
// - \&quot;微版本号\&quot;: 通过Http请求头X-OpenStack-Nova-API-Version来使用，从2.27版本后更改为OpenStack-API-Version。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) NovaShowVersion(request *model.NovaShowVersionRequest) (*model.NovaShowVersionResponse, error) {
	requestDef := GenReqDefForNovaShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NovaShowVersionResponse), nil
	}
}

// NovaShowVersionInvoker 查询指定API版本信息
func (c *EcsClient) NovaShowVersionInvoker(request *model.NovaShowVersionRequest) *NovaShowVersionInvoker {
	requestDef := GenReqDefForNovaShowVersion()
	return &NovaShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询任务的执行状态
//
// 查询Job的执行状态。
//
// 对于创建云服务器、删除云服务器、云服务器批量操作和网卡操作等异步API，命令下发后，会返回job_id，通过job_id可以查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询任务的执行状态
func (c *EcsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
