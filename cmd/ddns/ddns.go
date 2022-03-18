/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:00 CST
 * @Desc:
 */
package main

import (
	"github.com/AghostPrj/ddns/internal/initializator"
	"github.com/AghostPrj/ddns/internal/runtime"
)

func main() {
	initializator.InitApp()
	runtime.MainLoop()
}
