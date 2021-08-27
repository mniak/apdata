package apdata

import (
	"log"

	"github.com/HereMobilityDevelopers/mediary"
)

type (
	ClientOption      func(opt ClientOptionsData) ClientOptionsData
	ClientOptionsData struct {
		logger       *log.Logger
		interceptors []mediary.Interceptor
		debug        bool
	}
)

func defaultClientOptionsData() ClientOptionsData {
	return ClientOptionsData{
		logger:       log.Default(),
		interceptors: make([]mediary.Interceptor, 0),
	}
}

func combineOptions(options []ClientOption) ClientOptionsData {
	optionsData := defaultClientOptionsData()
	for _, opt := range options {
		optionsData = opt(optionsData)
	}
	return optionsData
}

func WithLogger(logger *log.Logger) ClientOption {
	return func(opt ClientOptionsData) ClientOptionsData {
		opt.logger = logger
		return opt
	}
}

func WithDebug(debug bool) ClientOption {
	return func(opt ClientOptionsData) ClientOptionsData {
		opt.debug = debug
		return opt
	}
}

func WithInterceptors(inter ...mediary.Interceptor) ClientOption {
	return func(opt ClientOptionsData) ClientOptionsData {
		opt.interceptors = append(opt.interceptors, inter...)
		return opt
	}
}
