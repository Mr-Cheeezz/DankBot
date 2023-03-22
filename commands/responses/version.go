package responses

import "github.com/mr-cheeezz/dankbot/env"

func Version() string {
	env.Load()
	return "DankBot - Go | " + env.Get("VERSION")
}
