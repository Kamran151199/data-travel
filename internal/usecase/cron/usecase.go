package cron

type UseCase interface {
	Start() error
}
