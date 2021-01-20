package mgo

type Option func(mgo *mgo)

func WithURI(uri string) Option {
	return func(mgo *mgo) {
		mgo.uri = uri
	}
}
