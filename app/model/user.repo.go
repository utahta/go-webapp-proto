package model

// Symfony でいう Repository 的な感じ
// User などのテーブルクラスは自動生成なので、触らないことという感じが良さそう
//
// 使い方はこうなりそう
//
// repo := model.NewUserRepository()
// repo.Find(1)
//

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Find(id int) (*User, error) {
	e, err := NewEngine()
	if err != nil {
		return nil, err
	}

	var user User
	if _, err := e.Id(id).Get(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
