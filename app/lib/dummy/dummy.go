package dummy

import (
	"github.com/utahta/go-webapp-proto/app/model"
)

type Dummy struct {
}

func New() *Dummy {
	return &Dummy{}
}

func (d *Dummy) Do(id int) (*model.User, error) {
	u, err := model.NewUserRepository().Find(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
