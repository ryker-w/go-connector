package mongodb

func (conn *Connection) Database(db string) *Database {
	database := &Database{
		db: conn.conn.Database(db),
	}
	return database
}

func (conn *Connection) Close() {
	_ = conn.conn.Disconnect(getContext())
}
