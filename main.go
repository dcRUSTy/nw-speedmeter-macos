package main

//#include <stdio.h>
//#include <ifaddrs.h>
//#include <sys/socket.h>
//#include <stdlib.h>
//#include <net/if_var.h>
//#include <string.h>
//long iBytes=0;
//long oBytes=0;
//void takeSnapshot(){
//    struct ifaddrs *ifa_list = 0, *ifa;
//    oBytes=0; iBytes=0;
//    getifaddrs(&ifa_list);
//        for (ifa = ifa_list; ifa; ifa = ifa->ifa_next) {
//            if (AF_LINK != ifa->ifa_addr->sa_family)
//                    continue;
//            if (ifa->ifa_data == 0)
//                continue;
//			  if((ifa->ifa_flags & 0x8)!=0)//IFF_LOOPBACK
//				continue;
//			  struct if_data *if_data = (struct if_data *)ifa->ifa_data;
//			  iBytes+=if_data->ifi_ibytes;
//			  oBytes+=if_data->ifi_obytes;
//        }
//	   free(ifa_list);
//}
//long getIBytes(){
//    return iBytes;
//}
//long getOBytes(){
//    return oBytes;
//}
import "C"
import (
	"fmt"
	"github.com/getlantern/systray"
	"time"
)

func main() {
	systray.Run(onReady, func() {})
}
func onReady() {
	systray.SetTitle("-/-kBps")
	for {
		thenTime := uint64(time.Now().UnixNano() / 1000000)
		thenIBytes, thenOBytes := getStats()
		time.Sleep(1000 * time.Millisecond)
		nowTime := uint64(time.Now().UnixNano() / 1000000)
		nowIBytes, nowOBytes := getStats()
		iSpeed := (float64(nowIBytes-thenIBytes) / 1024) / (float64(nowTime-thenTime) / 1000)
		oSpeed := (float64(nowOBytes-thenOBytes) / 1024) / (float64(nowTime-thenTime) / 1000)
		iSpeedString := fmt.Sprintf("%0.0f", oSpeed)
		oSpeedString := fmt.Sprintf("%0.0f", iSpeed)
		systray.SetTitle(oSpeedString + "/" + iSpeedString + "kBps")
	}
}

func getStats() (uint64, uint64) {
	C.takeSnapshot()
	pastIBytes := uint64(C.getIBytes())
	pastOBytes := uint64(C.getOBytes())
	return pastIBytes, pastOBytes
}
