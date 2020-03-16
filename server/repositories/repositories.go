package repositories

// Repository respresents the interface that must be implemented by
// any repository to be registered
type Repository interface {
	Find(query interface{}, options FindOptions) (interface{}, error)
	Insert(data interface{}) error
	Update(query interface{}, update interface{}) error
	Remove(query interface{}) error
}

// FindOptions represents options available to find methods implementing the interface
type FindOptions struct {
	Limit int      // limit of documents to return
	Sort  []string // fields to sort by
	Iter  bool     // if true, return an iterator instead of a slice of results
	Skip  int      // skip over n documents of the query result
}
