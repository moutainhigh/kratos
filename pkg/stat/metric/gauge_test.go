package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaugeAdd(t *testing.T) {
	gauge := NewGauge(GaugeOpts{})
	gauge.Add(100)
	gauge.Add(-50)
	val := gauge.Value()
	assert.Equal(t, val, int64(50))
}

func TestGaugeSet(t *testing.T) {
	gauge := NewGauge(GaugeOpts{})
	gauge.Add(100)
	gauge.Set(50)
	val := gauge.Value()
	assert.Equal(t, val, int64(50))
}

func TestGaugeVec(t *testing.T) {
	gaugeVec := NewGaugeVec(&GaugeVecOpts{
		Namespace: "test",
		Subsystem: "test",
		Name:      "test",
		Help:      "this is test metrics.",
		Labels:    []string{"name", "addr"},
	})
	gaugeVec.Set(float64(22.33), "name1", "127.0.0.1")
	assert.Panics(t, func() {
		NewGaugeVec(&GaugeVecOpts{
			Namespace: "test",
			Subsystem: "test",
			Name:      "test",
			Help:      "this is test metrics.",
			Labels:    []string{"name", "addr"},
		})
	}, "Expected to panic.")
	assert.NotPanics(t, func() {
		NewGaugeVec(&GaugeVecOpts{
			Namespace: "test",
			Subsystem: "test",
			Name:      "test2",
			Help:      "this is test metrics.",
			Labels:    []string{"name", "addr"},
		})
	}, "Expected normal. no panic.")
}
