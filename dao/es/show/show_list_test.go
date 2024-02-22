package show

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/iooikaak/microService2/database/es/show"

	"github.com/iooikaak/frame/util"
	"github.com/iooikaak/frame/xlog"
)

func TestDao_GetShowList(t *testing.T) {
	a, b := d.GetShowList(ctx, "ciqSDH8BmTViIP89hyje")
	t.Logf("%v---%v", a, b)
}

func TestDao_GetShowLists(t *testing.T) {
	a, b := d.GetShowLists(ctx, &show.ShowListSearchEsModel{
		From:     0,
		PageSize: 10,
	})
	var showLists []*show.ShowListEsModel
	for _, i := range a.Hits.Hits {
		item := new(show.ShowListEsModel)
		d := json.NewDecoder(bytes.NewReader(i.Source))
		d.UseNumber()
		err := d.Decode(item)
		if err != nil {
			xlog.Errorf("GetShowLists NewDecoder failed err:(%+v)", err)
			continue
		}

		showLists = append(showLists, item)
	}
	t.Logf("%v---%v---%#v", a, b, showLists)
}

func TestDao_InsertShowList(t *testing.T) {
	a := d.InsertShowList(ctx, &show.ShowListEsModel{
		StdShowId:            "602f95635f45b951eda10da8",
		BizShowId:            "ssssssssss",
		BizArtistIds:         "ddddddddd",
		IdentityRequiredType: 0,
		IsShowStdContent:     false,
		LastShowTime:         util.JsonTime{},
	})
	t.Logf("----%v---", a)
}

func TestDao_UpdateShowListByQueryAttr(t *testing.T) {
	a := d.UpdateShowListByQueryAttr(ctx, "602f95635f45b951eda10da8", time.Now())
	t.Logf("----%v---", a)
}
