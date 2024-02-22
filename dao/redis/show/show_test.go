package show

import (
	"testing"
	//"time"
	//conf "github.com/iooikaak/microService2/config"
)

//func TestDao_LockSupplier(t *testing.T) {
//	a := make(chan struct{})
//	r, err := d.LockShow(ctx, 114)
//	t.Logf("%v-----%v", r, err)
//	ti := time.NewTicker(10 * time.Second)
//	for i := range ti.C {
//		url, _ := conf.GetPxqUrl()
//		t.Logf("%#v-----%v---%v", i, conf.M.Get("pxq_url"), url)
//	}
//	<-a
//}

func TestDao_GetSupplierLock(t *testing.T) {
	r, err := d.GetShowLock(ctx, 111)
	defer func() {
		if r {
			t.Logf("fffff")
			_ = d.ReleaseShowLock(ctx, 111)
		}
	}()
	t.Logf("%v-----%v", r, err)
	if r {
		t.Logf("dddddd")
		_ = d.ReleaseShowLock(ctx, 111)
		r = false
	}
	t.Logf("%v-----%v", r, err)
}

func TestDao_ReleaseSupplierLock(t *testing.T) {
	err := d.ReleaseShowLock(ctx, 111)
	t.Logf("%v", err)
}
