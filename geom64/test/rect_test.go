package geomTest

import (
	"github.com/tadeuszjt/geom/geom64"
	"testing"
)

func rectIdentical(a, b geom.Rect) bool {
	return vec2Identical(a.Min, b.Min) && vec2Identical(a.Max, b.Max)
}

func TestRectZero(t *testing.T) {
	expected := geom.Rect{geom.Vec2{0, 0}, geom.Vec2{0, 0}}
	actual := geom.Rect{}
	if !rectIdentical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestRectOrigin(t *testing.T) {
	expected := geom.Rect{geom.Vec2{0, 0}, geom.Vec2{123, .456}}
	actual := geom.RectOrigin(123, .456)
	if !rectIdentical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestRectCentred(t *testing.T) {
	cases := []struct {
		width, height float64
		result        geom.Rect
	}{
		{0, 0, geom.Rect{}},
		{10, 20, geom.Rect{Min: geom.Vec2{-5, -10}, Max: geom.Vec2{5, 10}}},
		{12, 22, geom.Rect{Min: geom.Vec2{-6, -11}, Max: geom.Vec2{6, 11}}},
		{nan, pInf, geom.Rect{Min: geom.Vec2{nan, nInf}, Max: geom.Vec2{nan, pInf}}},
		{nInf, nan, geom.Rect{Min: geom.Vec2{pInf, nan}, Max: geom.Vec2{nInf, nan}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := geom.RectCentred(c.width, c.height)
		if !rectIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectCreate(t *testing.T) {
	cases := []struct {
		width, height float64
		position      geom.Vec2
		result        geom.Rect
	}{
		{0, 0, geom.Vec2{0, 0}, geom.Rect{}},
		{10, 20, geom.Vec2{0, 0}, geom.Rect{Min: geom.Vec2{-5, -10}, Max: geom.Vec2{5, 10}}},
		{10, 20, geom.Vec2{3, 4}, geom.Rect{Min: geom.Vec2{-2, -6}, Max: geom.Vec2{8, 14}}},
		{0, 0, geom.Vec2{3, 4}, geom.Rect{Min: geom.Vec2{3, 4}, Max: geom.Vec2{3, 4}}},
		{0.3, 0.8, geom.Vec2{-2.3, 4}, geom.Rect{Min: geom.Vec2{-2.45, 3.6}, Max: geom.Vec2{-2.15, 4.4}}},
		{-3, 0, geom.Vec2{1, 2}, geom.Rect{Min: geom.Vec2{2.5, 2}, Max: geom.Vec2{-0.5, 2}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := geom.RectCreate(c.width, c.height, c.position)
		if !rectIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectWidth(t *testing.T) {
	cases := []struct {
		rect  geom.Rect
		width float64
	}{
		{geom.Rect{Min: geom.Vec2{0, 0}, Max: geom.Vec2{0, 0}}, 0},
		{geom.Rect{Min: geom.Vec2{0, 0}, Max: geom.Vec2{10, 20}}, 10},
		{geom.Rect{Min: geom.Vec2{1.4, 3.2}, Max: geom.Vec2{2.3, 4.5}}, 0.9},
		{geom.Rect{Min: geom.Vec2{-8.2, 1.2}, Max: geom.Vec2{11.3, 4.5}}, 19.5},
	}

	for _, c := range cases {
		expected := c.width
		actual := c.rect.Width()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectHeight(t *testing.T) {
	cases := []struct {
		rect   geom.Rect
		height float64
	}{
		{geom.Rect{Min: geom.Vec2{0, 0}, Max: geom.Vec2{0, 0}}, 0},
		{geom.Rect{Min: geom.Vec2{0, 0}, Max: geom.Vec2{10, 20}}, 20},
		{geom.Rect{Min: geom.Vec2{1.4, 3.2}, Max: geom.Vec2{2.3, 4.5}}, 1.3},
		{geom.Rect{Min: geom.Vec2{8.2, -1.2}, Max: geom.Vec2{11.3, 4.5}}, 5.7},
	}

	for _, c := range cases {
		expected := c.height
		actual := c.rect.Height()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectContains(t *testing.T) {
	cases := []struct {
		rect    geom.Rect
		points  []geom.Vec2
		results []bool
	}{
		{
			geom.Rect{},
			[]geom.Vec2{{0, 0}, {0.00001, 0}, {0, 0.0001}, {-0.0001, 0}, {0, -0.00001}},
			[]bool{true, false, false, false, false},
		},
		{
			geom.Rect{Min: geom.Vec2{0, 0}, Max: geom.Vec2{10, 10}},
			[]geom.Vec2{{0, 0}, {-1, -1}, {1, -1}, {10, 10}, {10, 10.00001}},
			[]bool{true, false, false, true, false},
		},
		{
			geom.Rect{Min: geom.Vec2{-0.1, -0.2}, Max: geom.Vec2{0.1, 0.2}},
			[]geom.Vec2{{-0.1, -0.2}, {0.1, -0.21}, {0.1, 0.2}, {0, 0}, {nan, 0}},
			[]bool{true, false, true, true, false},
		},
		{
			geom.Rect{Min: geom.Vec2{1, 2}, Max: geom.Vec2{3, 4}},
			[]geom.Vec2{{1, 2}, {0.8, 1.9}, {2.8, 2.2}, {3.1, 0.9}, {1.1, 4}, {0.9, 4}},
			[]bool{true, false, true, false, true, false},
		},
		{
			geom.Rect{Min: geom.Vec2{100, 1.3}, Max: geom.Vec2{120, 1.8}},
			[]geom.Vec2{{110, 1.2}, {110, 1.3}, {110, 1.7}, {110, 1.9}},
			[]bool{false, true, true, false},
		},
	}

	for _, c := range cases {
		for i := range c.points {
			expected := c.results[i]
			actual := c.rect.Contains(c.points[i])
			if actual != expected {
				t.Errorf(
					"rect: %v, point: %v, expected: %v, got: %v",
					c.rect,
					c.points[i],
					expected,
					actual,
				)
			}
		}
	}
}

func TestRectVerts(t *testing.T) {
	cases := []struct {
		rect   geom.Rect
		result [4]geom.Vec2
	}{
		{
			geom.Rect{},
			[4]geom.Vec2{{0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		{
			geom.RectCentred(2, 2),
			[4]geom.Vec2{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}},
		},
		{
			geom.RectCreate(2, 2, geom.Vec2{-5, 5}),
			[4]geom.Vec2{{-6, 4}, {-4, 4}, {-4, 6}, {-6, 6}},
		},
		{
			geom.RectCreate(2, pInf, geom.Vec2{nInf, 5}),
			[4]geom.Vec2{{nInf, nInf}, {nInf, nInf}, {nInf, pInf}, {nInf, pInf}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.rect.Verts()
		
		for i := range expected {
			if !vec2Identical(expected[i], actual[i]) {
				t.Errorf("expected: %v, got: %v", expected, actual)
				break
			}
		}
	}
}