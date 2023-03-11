package utils

import "time"

type CommandCooldown struct {
	command  string
	lastUsed time.Time
}

func (c *CommandCooldown) Check(command string, cooldown time.Duration) bool {
	if command != c.command {
		return true
	}
	if time.Since(c.lastUsed) >= cooldown {
		return true
	}
	return false
}

func (c *CommandCooldown) Update(command string) {
	c.command = command
	c.lastUsed = time.Now()
}
