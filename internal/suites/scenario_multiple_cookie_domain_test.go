package suites

import (
	"context"
	"fmt"
	"log"
	"time"
)

// MultiCookieDomainScenario represents a set of tests for multi cookie domain suite.
type MultiCookieDomainScenario struct {
	*RodSuite

	domain, nextDomain string

	remember bool
}

// NewMultiCookieDomainScenario returns a new Multi Cookie Domain Test Scenario.
func NewMultiCookieDomainScenario(domain, nextDomain string, remember bool) *MultiCookieDomainScenario {
	return &MultiCookieDomainScenario{
		RodSuite:   new(RodSuite),
		domain:     domain,
		nextDomain: nextDomain,
		remember:   remember,
	}
}

func (s *MultiCookieDomainScenario) SetupSuite() {
	browser, err := StartRod()
	if err != nil {
		log.Fatal(err)
	}

	s.RodSession = browser

	err = updateDevEnvFileForDomain(s.domain, false)
	s.Require().NoError(err)
}

func (s *MultiCookieDomainScenario) TearDownSuite() {
	err := s.RodSession.Stop()

	if err != nil {
		log.Fatal(err)
	}
}

func (s *MultiCookieDomainScenario) SetupTest() {
	s.Page = s.doCreateTab(s.T(), HomeBaseURL)
	s.verifyIsHome(s.T(), s.Page)
}

func (s *MultiCookieDomainScenario) TearDownTest() {
	s.collectCoverage(s.Page)
	s.MustClose()
}

func (s *MultiCookieDomainScenario) TestRememberMe() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
		s.collectScreenshot(ctx.Err(), s.Page)
	}()

	s.doVisitLoginPage(s.T(), s.Page, s.domain, "")

	s.WaitElementLocatedByID(s.T(), s.Context(ctx), "username-textfield")

	has := s.CheckElementExistsLocatedByID(s.T(), s.Context(ctx), "remember-checkbox")

	s.Assert().Equal(s.remember, has)
}

func (s *MultiCookieDomainScenario) TestShouldAuthorizeSecret() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
		s.collectScreenshot(ctx.Err(), s.Page)
	}()

	targetURL := fmt.Sprintf("%s/secret.html", SingleFactorBaseURLFmt(s.domain))
	s.doLoginOneFactor(s.T(), s.Context(ctx), "john", "password", s.remember, s.domain, targetURL)
	s.verifySecretAuthorized(s.T(), s.Page)
}

func (s *MultiCookieDomainScenario) TestShouldRequestLoginOnNextDomainAfterLoginOnFirstDomain() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
		s.collectScreenshot(ctx.Err(), s.Page)
	}()

	firstDomainTargetURL := fmt.Sprintf("%s/secret.html", SingleFactorBaseURLFmt(s.domain))
	nextDomainTargetURL := fmt.Sprintf("%s/secret.html", SingleFactorBaseURLFmt(s.nextDomain))

	s.doLoginOneFactor(s.T(), s.Context(ctx), "john", "password", s.remember, s.domain, firstDomainTargetURL)
	s.verifySecretAuthorized(s.T(), s.Page)

	s.doVisit(s.T(), s.Page, nextDomainTargetURL)
	s.verifyIsFirstFactorPage(s.T(), s.Page)
}

func (s *MultiCookieDomainScenario) TestShouldStayLoggedInOnNextDomainWhenLoggedOffOnFirstDomain() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
		s.collectScreenshot(ctx.Err(), s.Page)
	}()

	firstDomainTargetURL := fmt.Sprintf("%s/secret.html", SingleFactorBaseURLFmt(s.domain))
	nextDomainTargetURL := fmt.Sprintf("%s/secret.html", SingleFactorBaseURLFmt(s.nextDomain))

	s.doLoginOneFactor(s.T(), s.Context(ctx), "john", "password", s.remember, s.domain, firstDomainTargetURL)
	s.verifySecretAuthorized(s.T(), s.Page)

	err := updateDevEnvFileForDomain(s.nextDomain, false)
	s.Require().NoError(err)

	s.doLoginOneFactor(s.T(), s.Context(ctx), "john", "password", !s.remember, s.nextDomain, nextDomainTargetURL)
	s.verifySecretAuthorized(s.T(), s.Page)

	s.doVisit(s.T(), s.Page, fmt.Sprintf("%s%s", GetLoginBaseURL(s.domain), "/logout"))
	s.verifyIsFirstFactorPage(s.T(), s.Page)

	s.doVisit(s.T(), s.Page, nextDomainTargetURL)
	s.verifySecretAuthorized(s.T(), s.Page)
}
