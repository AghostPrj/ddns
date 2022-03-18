/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package runtime

import (
	"github.com/AghostPrj/ddns/internal/initializator"
	"github.com/AghostPrj/ddns/internal/utils/ipUtils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func MainLoop() {
	for {
		result, err := ipUtils.GetLocalIpByInterfaceName(viper.GetString(initializator.ConfUpstreamInterfaceNameKey))
		result = &ipUtils.IpAddrPayload{
			V4: make([]string, 1),
			V6: make([]string, 1),
		}
		err = nil
		result.V6[0] = "240e:368:c0f:cf60:a84a:17fa:ba4a:fbe5"
		result.V4[0] = "59.172.116.189"
		if err != nil {
			log.WithField("err", err).Error()
		} else {
			log.WithField("data", result).Info()
		}
		time.Sleep(time.Second * viper.GetDuration(initializator.ConfAppLoopDelayKey))
	}
}
