package cmd

import "testing"

func TestSecurityContextNilRORF(t *testing.T) {
	runAuditTest(t, "security_context_nil.yml", auditReadOnlyRootFS, []int{ErrorReadOnlyRootFilesystemNil})
}

func TestReadOnlyRootFilesystemNil(t *testing.T) {
	runAuditTest(t, "read_only_root_filesystem_nil.yml", auditReadOnlyRootFS, []int{ErrorReadOnlyRootFilesystemNil})
}

func TestReadOnlyRootFilesystemFalse(t *testing.T) {
	runAuditTest(t, "read_only_root_filesystem_false.yml", auditReadOnlyRootFS, []int{ErrorReadOnlyRootFilesystemFalse})
}

func TestReadOnlyRootFilesystemFalseAllowed(t *testing.T) {
	runAuditTest(t, "read_only_root_filesystem_false_allowed.yml", auditReadOnlyRootFS, []int{ErrorReadOnlyRootFilesystemFalseAllowed})
}

func TestReadOnlyRootFilesystemMisconfiguredAllow(t *testing.T) {
	runAuditTest(t, "read_only_root_filesystem_misconfigured_allow.yml", auditReadOnlyRootFS, []int{ErrorMisconfiguredKubeauditAllow})
}
