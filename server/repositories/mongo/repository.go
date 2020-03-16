package repositories

import (
	"github.com/edwintcloud/print-solution/server/repositories"
	"github.com/globalsign/mgo"
)

type mongoRepository struct {
	c *mgo.Collection
}

// NewMongoRepository creates a new mongo repo
func NewMongoRepository(session *mgo.Session, dbName string, colName string) repositories.Repository {
	return &mongoRepository{
		c: session.DB(dbName).C(colName),
	}

}

func (r *mongoRepository) Find(query interface{}, options repositories.FindOptions) (interface{}, error) {
	var result []map[string]interface{}

	findQuery := r.c.Find(query)
	if options.Limit > 0 {
		findQuery = findQuery.Limit(options.Limit)
	}
	if options.Skip > 0 {
		findQuery = findQuery.Skip(options.Skip)
	}
	if len(options.Sort) > 0 {
		findQuery = findQuery.Sort(options.Sort...)
	}

	if options.Iter {
		return findQuery.Iter(), nil
	}

	err := findQuery.All(&result)

	return result, err
}

func (r *mongoRepository) Insert(data interface{}) error {
	return r.c.Insert(data)
}

func (r *mongoRepository) Update(query interface{}, update interface{}) error {
	return r.c.Update(query, update)
}

func (r *mongoRepository) Remove(query interface{}) error {
	return r.c.Remove(query)
}
