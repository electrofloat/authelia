package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func NewMultiCookieDomainSuite() *MultiCookieDomainSuite {
	return &MultiCookieDomainSuite{}
}

type MultiCookieDomainSuite struct {
	suite.Suite
}

func (s *MultiCookieDomainSuite) TestMultiCookieDomainFirstDomainScenario() {
	suite.Run(s.T(), NewMultiCookieDomainScenario(BaseDomain, Example2DotCom, "authelia_session", true))
}

func (s *MultiCookieDomainSuite) TestMultiCookieDomainSecondDomainScenario() {
	suite.Run(s.T(), NewMultiCookieDomainScenario(Example2DotCom, BaseDomain, "example2_session", false))
}

func (s *MultiCookieDomainSuite) TestMultiCookieDomainThirdDomainScenario() {
	suite.Run(s.T(), NewMultiCookieDomainScenario(Example3DotCom, BaseDomain, "authelia_session", true))
}

func TestMultiCookieDomainSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewMultiCookieDomainSuite())
}
