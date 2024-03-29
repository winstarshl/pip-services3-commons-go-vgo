package commands

import (
	"github.com/winstarshl/pip-services3-commons-go-vgo/run"
	"github.com/winstarshl/pip-services3-commons-go-vgo/validate"
)

type InterceptedCommand struct {
	interceptor ICommandInterceptor
	next        ICommand
}

func NewInterceptedCommand(interceptor ICommandInterceptor, next ICommand) *InterceptedCommand {
	return &InterceptedCommand{
		interceptor: interceptor,
		next:        next,
	}
}

func (c *InterceptedCommand) Name() string {
	return c.interceptor.Name(c.next)
}

func (c *InterceptedCommand) Execute(correlationId string, args *run.Parameters) (result interface{}, err error) {
	return c.interceptor.Execute(correlationId, c.next, args)
}

func (c *InterceptedCommand) Validate(args *run.Parameters) []*validate.ValidationResult {
	return c.interceptor.Validate(c.next, args)
}
