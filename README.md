# Growth 2021 H1

## Test Double

Gerard Meszaros introduced the differences between test types in his
_xUnit Test Patterns_ book published in May 2007.

The generic term he used to define a simplified version of an object or
procedure to reduce complexity and ease testing is _Test Double_.

From this generic terms come several more precise terms.

The [SUT](/service/service.go) _(System Under Test)_ is a service on which
we inject doubles to test interactions with its datastore.

## Dummy

Dummy objects are being passed around without being used.

## Fake

Fake objects have simplified implementations.

## Stubs

Stubs are limited to what is defined for the test. It doesn't have
expectations.

## Spies

Spies are stubs which also record usage. Expectations are evaluated after
the execution of the test.

## Mocks

Mocks define expectations before the test is executed and may throw errors
if an unexpected call was made.
