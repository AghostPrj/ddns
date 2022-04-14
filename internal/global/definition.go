/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package global

import "github.com/AghostPrj/ddns/internal/utils/dnsUtils"

const (
	DomainServiceProviderAliyun = "aliyun"
	DomainServiceProviderDnspod = "dnspod"
)

const (
	ApplicationName = "ddns"

	ConfAppLoopDelayKey    = "app.loop.delay"
	EnvAppLoopDelayKey     = "app_loop_delay"
	DefaultAppLoopDelayKey = 60

	ConfDomainServiceProvider = "app.domain.provider"
	EnvDomainServiceProvider  = "app_domain_provider"

	ConfDnspodTokenSecretKey = "app.token.dnspod.secret"
	EnvDnspodTokenSecretKey  = "app_token_dnspod_secret"

	ConfAliyunTokenIdKey = "app.token.aliyun.id"
	EnvAliyunTokenIdKey  = "app_token_aliyun_id"

	ConfAliyunTokenSecretKey = "app.token.aliyun.secret"
	EnvAliyunTokenSecretKey  = "app_token_aliyun_secret"

	ConfDomainKey = "app.domain"
	EnvDomainKey  = "app_domain"

	ConfSubDomainKey = "app.domain.sub"
	EnvSubDomainKey  = "app_domain_sub"

	ConfUpstreamInterfaceNameKey = "app.interface.name"
	EnvUpstreamInterfaceNameKey  = "app_interface_name"
	DefaultUpstreamInterfaceName = "pppoe-wan"

	DomainTypeIpv4Direct = "A"
	DomainTypeIpv6Direct = "AAAA"
)

var (
	DescribeRecordFunction     func() (*dnsUtils.DomainRecordPayload, error)
	UpdateDomainRecordFunction func(*string, *string, *string) error
	AddDomainRecordFunction    func(*string, *string) error
)
