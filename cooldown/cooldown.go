package cooldown

import (
	"sync"
	"time"
)

type Cooldowns struct {
	mu        sync.Mutex
	cooldowns map[string]time.Time
}

func NewCooldowns() *Cooldowns {
	return &Cooldowns{
		cooldowns: make(map[string]time.Time),
	}
}

func (c *Cooldowns) OnCooldown(command string, duration time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	lastUsed, ok := c.cooldowns[command]
	if !ok || time.Since(lastUsed) >= duration {
		c.cooldowns[command] = time.Now()
		return false
	}
	return true
}
