package godatastruct

import (
  . "launchpad.net/gocheck"
  "testing"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}
var _ = Suite(&MySuite{})

func setTests(c *C, s Set) {
  c.Assert(s.Set(1), Equals, true)
  c.Assert(s.Exists(1), Equals, true)
  c.Assert(s.Exists(2), Equals, false)
  c.Assert(s.Len(), Equals, 1)
  s.Clear()

  c.Assert(s.Len(), Equals, 0)
  c.Assert(s.Exists(1), Equals, false)

  c.Assert(s.Set(1), Equals, true)
  c.Assert(s.Set("Cats"), Equals, true)
  c.Assert(s.Len(), Equals, 2)
  c.Assert(s.Erase("Cats"), Equals, true)
  c.Assert(s.Erase("Dogs"), Equals, false)
  c.Assert(s.Len(), Equals, 1)
}

func (s *MySuite) TestUnorderedSet(c *C) {
  us := NewUnorderedSet() 
  setTests(c, us)
}

func (s *MySuite) TestThreadSafeUnorderedSet(c *C) {
  us:= NewTheadSafeUnorderedSet() 
  setTests(c, us)
}
