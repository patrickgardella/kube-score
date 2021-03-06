package score

import (
	"testing"
)

func TestIngressTargetsService(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "ingress-targets-service.yaml", "Ingress targets Service", 10)
}

func TestIngressTargetsServiceNoMatch(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "ingress-targets-service-no-match.yaml", "Ingress targets Service", 1)
}
