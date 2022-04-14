/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/ddns/internal/global"
	"github.com/AghostPrj/ddns/internal/utils/aliyunDnsUtils"
	"github.com/AghostPrj/ddns/internal/utils/dnspodUtils"
	"github.com/AghostPrj/ddns/internal/utils/ipUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/configUtils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitApp() {
	configUtils.SetConfigFileName(global.ApplicationName)
	bindAppConfigKey()
	bindAppConfigDefaultValue()
	configUtils.InitConfigAndLog()
	checkAppConfig()
	ipUtils.InitNetLinkConn()
	switch viper.GetString(global.ConfDomainServiceProvider) {
	case global.DomainServiceProviderAliyun:
		aliyunDnsUtils.InitDnsClient()
		global.DescribeRecordFunction = aliyunDnsUtils.DescribeRecord
		global.UpdateDomainRecordFunction = aliyunDnsUtils.UpdateDomainRecord
		global.AddDomainRecordFunction = aliyunDnsUtils.AddDomainRecord
		break
	case global.DomainServiceProviderDnspod:
		dnspodUtils.InitDnsClient()
		global.DescribeRecordFunction = dnspodUtils.DescribeRecord
		global.AddDomainRecordFunction = dnspodUtils.AddDomainRecord
		global.UpdateDomainRecordFunction = dnspodUtils.UpdateDomainRecord

		break
	}
}

func bindAppConfigKey() {
	configUtils.ConfigKeyList = append(configUtils.ConfigKeyList,
		[]string{global.ConfDomainServiceProvider, global.EnvDomainServiceProvider},
		[]string{global.ConfDnspodTokenSecretKey, global.EnvDnspodTokenSecretKey},
		[]string{global.ConfAliyunTokenIdKey, global.EnvAliyunTokenIdKey},
		[]string{global.ConfAliyunTokenSecretKey, global.EnvAliyunTokenSecretKey},
		[]string{global.ConfAppLoopDelayKey, global.EnvAppLoopDelayKey},
		[]string{global.ConfDomainKey, global.EnvDomainKey},
		[]string{global.ConfSubDomainKey, global.EnvSubDomainKey},
		[]string{global.ConfUpstreamInterfaceNameKey, global.EnvUpstreamInterfaceNameKey},
	)
}
func bindAppConfigDefaultValue() {
	viper.SetDefault(global.ConfAppLoopDelayKey, global.DefaultAppLoopDelayKey)
	viper.SetDefault(global.ConfUpstreamInterfaceNameKey, global.DefaultUpstreamInterfaceName)
}

func checkAppConfig() {
	if !checkDomainServiceProviderConf() {
		log.WithField("err", "domain provider config error").
			WithField("op", "init").
			Panic("config error")
	}

	if !checkDomainConfig() {
		log.WithField("err", "domain error").
			WithField("op", "init").
			Panic("config error")
	}
}
