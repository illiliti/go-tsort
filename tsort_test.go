package tsort

import (
	"errors"
	"testing"
)

func TestCircularLoop(t *testing.T) {
	ss := []struct {
		test     string
		sample   []string
		expected error
	}{
		{
			test: "BAD:pkg_1->pkg_2->pkg_1",
			sample: []string{
				"pkg_1", "pkg_2",
				"pkg_2", "pkg_1",
			},
			expected: ErrCircularLoop,
		},
		{
			test: "BAD:pkg_1->pkg_2->pkg_3->pkg_1",
			sample: []string{
				"pkg_1", "pkg_2",
				"pkg_2", "pkg_3",
				"pkg_3", "pkg_1",
			},
			expected: ErrCircularLoop,
		},
		{
			test: "GOOD:pkg_1->pkg_2->pkg_3",
			sample: []string{
				"pkg_1", "pkg_2",
				"pkg_2", "pkg_3",
			},
			expected: nil,
		},
	}

	for _, s := range ss {
		s := s // HACK

		t.Run(s.test, func(t *testing.T) {
			t.Parallel()

			var (
				tg Graph
				vx *Vertex
			)

			for _, n := range s.sample {
				switch {
				case vx == nil:
					vx = tg.AddVertex(n)
				case vx.Name != n:
					vx.AddEdge(n)
					vx = nil
				default:
					vx = nil
				}
			}

			_, err := tg.Sort()

			if !errors.Is(err, s.expected) {
				t.Errorf("got %q, want %q", err, s.expected)
			}
		})
	}
}
