package datastore

import "gopkg.in/mgo.v2"

type GridFSDataStore struct {
	session *mgo.Session
	dbName  string
	prefix  string
}

func NewGridFSDataStore(session *mgo.Session, dbName string, prefix string) *GridFSDataStore {
	return &GridFSDataStore{
		session: session,
		dbName:  dbName,
		prefix:  prefix,
	}
}

func (ds *GridFSDataStore) copy() *GridFSDataStore {
	return NewGridFSDataStore(ds.session.Copy(), ds.dbName, ds.prefix)
}

func (ds *GridFSDataStore) db() *mgo.Database {
	return ds.copy().session.DB(ds.dbName)
}

func (ds *GridFSDataStore) gridFS() *mgo.GridFS {
	return ds.db().GridFS(ds.prefix)
}

func (ds *GridFSDataStore) FindByName(name string) (*mgo.GridFile, error) {
	return ds.gridFS().Open(name)
}

func (ds *GridFSDataStore) Create(name string) (*mgo.GridFile, error) {
	return ds.gridFS().Create(name)
}

func (ds *GridFSDataStore) ListFileNames() ([]string, error) {
	var result []string
	err := ds.gridFS().Find(nil).Distinct("filename", &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ds *GridFSDataStore) Close() {
	ds.session.Close()
}
