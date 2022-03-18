/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/ddns/internal/utils/aliyunDnsUtils"
	"github.com/AghostPrj/ddns/internal/utils/ipUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/configUtils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitApp() {
	configUtils.SetConfigFileName(ApplicationName)
	bindAppConfigKey()
	bindAppConfigDefaultValue()
	configUtils.InitConfigAndLog()
	checkAppConfig()
	ipUtils.InitNetLinkConn()
	aliyunDnsUtils.InitAliyunDnsClient()
}

func bindAppConfigKey() {
	configUtils.ConfigKeyList = append(configUtils.ConfigKeyList,
		[]string{ConfAliyunTokenIdKey, EnvAliyunTokenIdKey},
		[]string{ConfAliyunTokenSecretKey, EnvAliyunTokenSecretKey},
		[]string{ConfAppLoopDelayKey, EnvAppLoopDelayKey},
		[]string{ConfDomainKey, EnvDomainKey},
		[]string{ConfSubDomainKey, EnvSubDomainKey},
		[]string{ConfUpstreamInterfaceNameKey, EnvUpstreamInterfaceNameKey},
	)
}
func bindAppConfigDefaultValue() {
	viper.SetDefault(ConfAppLoopDelayKey, DefaultAppLoopDelayKey)
	viper.SetDefault(ConfUpstreamInterfaceNameKey, DefaultUpstreamInterfaceName)
}

func checkAppConfig() {
	if !checkAliyunToken() {
		log.WithField("err", "aliyun token error").
			WithField("op", "init").
			Panic("config error")
	}

	if !checkDomainConfig() {
		log.WithField("err", "domain error").
			WithField("op", "init").
			Panic("config error")
	}
}
