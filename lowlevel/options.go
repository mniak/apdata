package lowlevel

import (
	"github.com/HereMobilityDevelopers/mediary"
)

type (
	ClientOption      func(opt ClientOptionsData) ClientOptionsData
	ClientOptionsData struct {
		interceptors []mediary.Interceptor
		debug        bool
	}
)

func defaultClientOptionsData() ClientOptionsData {
	return ClientOptionsData{
		interceptors: []mediary.Interceptor{
			normalizeResponseJSON,
		},
	}
}

func combineOptions(options []ClientOption) ClientOptionsData {
	optionsData := defaultClientOptionsData()
	for _, opt := range options {
		optionsData = opt(optionsData)
	}
	return optionsData
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
