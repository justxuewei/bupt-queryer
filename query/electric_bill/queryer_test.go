package electric_bill

import (
	"github.com/xavier-niu/bupt-queryer/request"
	"testing"
)

func TestQueryer_Query(t *testing.T) {
	session := request.NewSession("32341", "grgg")
	queryer := NewElectricBillQueryer(session)
	remaining, free, err := queryer.Query("gr", "5", "102-1638", "1")
	if err != nil {
		t.Error(err)
	}
	t.Log(remaining, free)
}
