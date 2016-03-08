package health

import "testing"

func TestNewHealth(t *testing.T) {
	h := NewHealth()

	if !h.IsDown() {
		t.Errorf("NewHealth().IsDown() == %t, want %t", h.IsDown(), true)
	}
}

func Test_Health_Unknown(t *testing.T) {
	h := NewHealth()
	h.Unknown()

	if !h.IsUnknown() {
		t.Errorf("NewHealth().IsUnknown() == %t, want %t", h.IsUnknown(), true)
	}
}

func Test_Health_Up(t *testing.T) {
	h := NewHealth()
	h.Up()

	if !h.IsUp() {
		t.Errorf("NewHealth().IsUp() == %t, want %t", h.IsUp(), true)
	}
}

func Test_Health_Down(t *testing.T) {
	h := NewHealth()
	h.Up()
	h.Down()

	if !h.IsDown() {
		t.Errorf("NewHealth().IsDown() == %t, want %t", h.IsDown(), true)
	}
}

func Test_Health_OutOfService(t *testing.T) {
	h := NewHealth()
	h.OutOfService()

	if !h.IsOutOfService() {
		t.Errorf("NewHealth().IsOutOfService() == %t, want %t", h.IsOutOfService(), true)
	}
}

func Test_Health_IsUp(t *testing.T) {
	h := NewHealth()
	h.Up()

	if h.Status != Up {
		t.Errorf("NewHealth().Status == %t, want %t", h.Status, Up)
	}
}

func Test_Health_IsDown(t *testing.T) {
	h := NewHealth()

	if h.Status != Down {
		t.Errorf("NewHealth().Status == %t, want %t", h.Status, Down)
	}
}

func Test_Health_IsOutOfService(t *testing.T) {
	h := NewHealth()
	h.OutOfService()

	if h.Status != OutOfService {
		t.Errorf("NewHealth().Status == %t, want %t", h.Status, OutOfService)
	}
}
