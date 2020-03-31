package kf

import "kfonline/util/request"

func Login(r *request.Request) bool {
	if err := r.Validate([]string{"phone", "password"}); err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

func Uuid(r *request.Request) bool {
	return r.Success(nil)
}

func Publish(r *request.Request) bool {
	return r.Success(nil)
}

func History(r *request.Request) bool {
	return r.Success(nil)
}
