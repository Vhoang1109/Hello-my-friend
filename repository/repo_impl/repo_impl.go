package repo_impl

import (
	"context"
	"time"

	"example.com/test/model"
	"example.com/test/model/response"
	"example.com/test/repository"
)

type RepoImpl struct {
}

func NewRepo() repository.Repo {
	return &RepoImpl{}
}

func (r *RepoImpl) Process(context context.Context, input model.Input) (response.Output, error) {
	name := input.Name
	year := time.Now().Year() - input.Year
	output := response.Output{}
	output.Name = name
	output.Age = year
	return output, nil
}
