package configs

import "context"

type MyContext struct {
	Ctx context.Context
}

var appCtx *MyContext


func LoadContext() {
	appCtx = &MyContext{
		Ctx: context.Background(),
	}
}


func GetContext() context.Context {
	
	return appCtx.Ctx

}