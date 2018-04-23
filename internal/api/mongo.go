package api

import (
	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func init() {
	var err error
	mgoSession, err = mgo.Dial("11.11.1.6")
	if err != nil {
		panic(err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)
}
