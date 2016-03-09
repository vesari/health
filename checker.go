package health

// Checker is a interface used to provide an indication of application health.
type Checker interface {
	Check() Health
}

// CompositeChecker aggregate a list of Checkers
type CompositeChecker struct {
	checkers map[string]Checker
}

// NewCompositeChecker creates a new CompositeChecker
func NewCompositeChecker() CompositeChecker {
	return CompositeChecker{
		checkers: map[string]Checker{},
	}
}

// AddChecker add a Checker to the aggregator
func (c *CompositeChecker) AddChecker(name string, checker Checker) {
	c.checkers[name] = checker
}

// Check returns the combination of all checkers added
// if some check is not up, the combined is marked as down
func (c CompositeChecker) Check() Health {
	health := NewHealth()
	health.Up()

	healths := map[string]Health{}

	for name, check := range c.checkers {
		h := check.Check()

		if !health.IsDown() && h.IsDown() {
			health.Down()
		}

		healths[name] = h
	}

	health.Info = healths

	return health
}
