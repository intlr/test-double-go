# Test Double

Gerard Meszaros introduced the differences between test objects in his
_xUnit Test Patterns_ book published in May 2007.

The generic term he used to define a simplified version of an object or
procedure to reduce complexity and ease testing is _Test Double_.

From this generic terms come several more precise terms.

The [SUT](/service/service.go) _(System Under Test)_ is a service on which
we inject doubles to test interactions with its datastore.

## Dummy

[Dummys](/dummy/service_test.go) are being passed around without being
used.

## Fake

[Fakes](/fake/service_test.go) have simplified implementations.

## Stub

[Stubs](/stub/service_test.go) are limited to what is defined for the test,
they don't have expectations.

## Spy

[Spies](/spy/service_test.go) are stubs recording usage. Expectations are
evaluated after the execution of the test.

## Mock

[Mocks](/mock/service_test.go) define expectations before the test is
executed, throwing errors if an unexpected call was made.
