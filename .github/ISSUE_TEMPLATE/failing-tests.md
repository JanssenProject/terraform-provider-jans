---
name: Failing Test / Acceptance Test Failure
about: Report a failing acceptance test in the Janssen Terraform provider
title: '[TEST] '
labels: test-failure, terraform
assignees: ''

---

## Failing Test


### Description
Describe the failing test and its purpose (e.g., "Test for creating a new client fails with 401").


### Test Details
- **Test Name**: (e.g., `TestAccJansClient_basic`)
- **Test File/Location**: (e.g., `jans/resource_jans_client_test.go`)


### Steps to Reproduce
1. Run `make testacc TESTARGS='-run=TestAccJansClient'`
2. Observe failure


### Expected Behavior
Describe what the test should do when passing.


### Actual Behavior
Include error message/stack trace.


### Environment
- **Provider Version**: (e.g., v0.1.0)
- **Janssen Server Version**: (e.g., v1.2.0)
- **OS**: (e.g., Linux)
- **Terraform Version**: (e.g., v1.9.0)


### Logs/Screenshots
Attach test output or screenshots of failure.


### Additional Context
Any recent changes or related issues?
