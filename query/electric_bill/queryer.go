package electric_bill

import (
	"encoding/json"
	"errors"
	"github.com/xavier-niu/bupt-queryer/query"
	"github.com/xavier-niu/bupt-queryer/request"
	"io/ioutil"
	"net/url"
)

const queryUrl = "https://app.bupt.edu.cn/buptdf/wap/default/search"

type Queryer struct {
	query.AppQueryer
}

type info struct {
	request.BaseInfo
	D infoD `json:"d"`
}

type infoD struct {
	Data InfoData `json:"data"`
}

type InfoData struct {
	Surplus float32 `json:"surplus"`
	FreeEnd float32 `json:"freeEnd"`
}

func NewElectricBillQueryer(session *request.Session) *Queryer {
	return &Queryer{
		AppQueryer: query.NewAppQuery(session),
	}
}

func (q *Queryer) Query(apartmentId, floorId, dormitoryNumber, areaId string) (remaining float32, free float32, err error) {
	rsp, err := q.Session.PostForm(queryUrl, url.Values{
		"partmentId": {apartmentId},
		"floorId": {floorId},
		"dromNumber": {dormitoryNumber},
		"areaid": {areaId},
	})
	if err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	var info info
	err = json.Unmarshal(body, &info)
	if err != nil {
		return
	}
	if info.Code != 0 {
		err = errors.New(info.Message)
		return
	}
	remaining, free = info.D.Data.Surplus, info.D.Data.FreeEnd
	return
}
