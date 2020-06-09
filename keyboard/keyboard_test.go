package keyboard

import (
	"testing"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/stretchr/testify/assert"
)

// FIXME: Test cases are not portable
//   The hardcoded string representation/keysym associated to each keycode
//   is not portable across devices (due to different keyboard mappings).

func TestLookupString(t *testing.T) {
	InitX()

	// test input 'A'
	var event1 xproto.KeyPressEvent
	event1.State = uint16(0)
	event1.Detail = xproto.Keycode(38)

	ch1, _ := LookupString(&event1)
	assert.Equal(t, "a", ch1)

	// test input 'Shift + A'
	var event2 xproto.KeyPressEvent
	event2.State = uint16(1)
	event2.Detail = xproto.Keycode(38)

	ch2, _ := LookupString(&event2)
	assert.Equal(t, "A", ch2)
}

func TestLookupStringKeysym(t *testing.T) {
	InitX()

	// test input 'A'
	var event1 xproto.KeyPressEvent
	event1.State = uint16(0)
	event1.Detail = xproto.Keycode(38)

	_, keysym1 := LookupString(&event1)
	assert.Equal(t, uint32(0x0061), uint32(keysym1))

	// test input 'Shift + A'
	var event2 xproto.KeyPressEvent
	event2.State = uint16(1)
	event2.Detail = xproto.Keycode(38)

	_, keysym2 := LookupString(&event2)
	assert.Equal(t, uint32(0x0041), uint32(keysym2))

	// test input 'Return'
	var event3 xproto.KeyPressEvent
	event3.State = uint16(0)
	event3.Detail = xproto.Keycode(36)

	_, keysym3 := LookupString(&event3)
	assert.Equal(t, uint32(0xFF0D), uint32(keysym3))
}
