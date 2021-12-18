package employee

import "testing"

/*
	Test Driven Development (or Test First Development)  is a software  development approach where we write the tests
	before writing the code to have a clear and unbiased understanding of what the code SHOULD do.

	Writing the code first and test later leads to development of a biased opinion on what the code should do;
	hence, we must write the tests first to avoid multiple iterations on the code.
*/

/*
	- Any test file in go must end with `_test.go`. This is how go recognises that the file ignored
	while building the package.

	- Each test function must be prefixed by `Test` and the signature for testing functions must be:
	func TestXxx(*testing.T){
		// test your code here
	}

	- Test function naming conventions:
	To test a function named `Xyx` or `xyz`, the test name should be `TestXyx` (notice the Mixed case).
	For a more specific unit test, we can add suffix to differentiate between the tests, like:
	`TestXyzInvalidParameter`
	`TestXyzSuccess`
	`TestXyzError`
	While running the tests for entire project, it is preferable to differentiate between the tests from
	different packages:
	`TestEmployee_XyzSuccess` // student package
	(notice that the package name is followed by an `_`)

	- Any new helper function or a custom type created for testing purpose should be added to the test file ONLY
	to avoid being included in the package build.

	- The source code file and its corresponding test file(s) must be in the same package.

	--------------------------------------------------------------------------------------------------------------------
	NOTE:
	We can keep test files in a different package, but that approach is used for Black-box Testing,
	which allows us to test only the exported functions/methods of a package, like for the `fmt` package
	in go (not needed for now).

	The white box approach, where the test and the production code are in the same package,
	allows us to test both the non-exported and exported identifiers of the package.
	This is the preferable approach when writing unit tests that require access to non-exported
	variables, functions, and methods.

*/

func TestEmployee_Create(t *testing.T) {

}

func TestEmployee_GetAll(t *testing.T) {

}

func TestEmployee_Get(t *testing.T) {

}

func TestEmployee_Update(t *testing.T) {

}

func TestEmployee_Delete(t *testing.T) {

}
