package mongodb

func (db *Database) Collection(col string) *Collection {
	collection := &Collection{
		delegate: db.db.Collection(col),
	}
	return collection
}
