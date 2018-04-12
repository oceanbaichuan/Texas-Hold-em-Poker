package internal

import (
	"github.com/dolotech/leaf/db/mongodb"
	"github.com/golang/glog"
	"server/conf"
)



var mongoDB *mongodb.DialContext

const DB  = "mahjongdb"

func init()  {

	if conf.Server.DBMaxConnNum <= 0 {
		conf.Server.DBMaxConnNum = 100
	}
	db, err := mongodb.Dial(conf.Server.DBUrl, conf.Server.DBMaxConnNum)
	if err != nil {
		glog.Fatalf("dial mongodb error: %v", err)
	}
	mongoDB = db

	err = db.EnsureCounter(DB, "counters", "users")
	if err != nil {
		glog.Fatalf("ensure counter error: %v", err)
	}

	err = db.EnsureCounter(DB, "counters", "rooms")
	if err != nil {
		glog.Fatalf("ensure counter error: %v", err)
	}
}

func mongoDBDestroy()  {
	mongoDB.Close()
	mongoDB = nil

}

func mongoDBNextSeq(id string) (int, error) {
	return mongoDB.NextSeq(DB, "counters", id)
}
