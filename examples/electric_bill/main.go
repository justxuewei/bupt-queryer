package main

import (
	"flag"
	"fmt"
	"github.com/xavier-niu/bupt-queryer/notifier/server_chan"
	"github.com/xavier-niu/bupt-queryer/query/electric_bill"
	"github.com/xavier-niu/bupt-queryer/request"
	"log"
)

var (
	usr = flag.String("u", "", "username")
	pwd = flag.String("p", "", "password")
	apartId = flag.String("a", "", "apartment ID")
	fid = flag.String("f", "", "floor ID")
	dnum = flag.String("d", "", "dormitory number")
	areaId = flag.String("A", "", "area ID, xitucheng->1, shahe->2")
	sk = flag.String("s", "", "server_chan send key")
	threshold = flag.Float64("t", 0, "warning threshold, 0 means no limit")
)

func main() {
	flag.Parse()

	q := electric_bill.NewElectricBillQueryer(request.NewSession(*usr, *pwd))
	r, f, err := q.Query(*apartId, *fid, *dnum, *areaId)
	if err != nil {
		log.Panicf("electric bill query is failed: %v", err)
	}
	if r+f < float32(*threshold) {
		log.Printf("stop to dispatch message because remaining is greater than threshold")
		return
	}
	n := server_chan.NewNotifier(*sk)
	err = n.Notify("电费查询", fmt.Sprintf("剩余电费：%.2f, 赠送剩余电费：%.2f, 总计剩余电费：%.2f", r, f, r+f))
	if err != nil {
		log.Panicf("notifing via server chan is failed: %v", err)
	}
}
