/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package initializator

const (
	ApplicationName = "ddns"

	ConfAppLoopDelayKey    = "app.loop.delay"
	EnvAppLoopDelayKey     = "app_loop_delay"
	DefaultAppLoopDelayKey = 60

	ConfAliyunTokenIdKey = "app.token.aliyun.token.id"
	EnvAliyunTokenIdKey  = "app_token_aliyun_token_id"

	ConfAliyunTokenSecretKey = "app.token.aliyun.token.secret"
	EnvAliyunTokenSecretKey  = "app_token_aliyun_token_secret"

	ConfDomainKey = "app.domain"
	EnvDomainKey  = "app_domain"

	ConfSubDomainKey = "app.domain.sub"
	EnvSubDomainKey  = "app_domain_sub"

	ConfUpstreamInterfaceNameKey = "app.interface.name"
	EnvUpstreamInterfaceNameKey  = "app_interface_name"
	DefaultUpstreamInterfaceName = "pppoe-wan"
)
