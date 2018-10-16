# complist
package for comparison of two items

The package only has one function IsSame(interface{}, interface{}) (bool, error) which takes in two items and compares whether they have the same number type of elements. For now it works for all ints, uints, strings, slices, arrays, and maps. It DOES NOT work for comparing individual floats or structs (yet). 

IsSame compares two items according to the number and value of their elements. The two items are found to be identical if they have the same number of elements of the same value. So, "home" and "ohme", 12345 and 24153, a brand new box of Legos and another one already assembled would all be found identical.

Errors are returned if the two items are not of the same type or if the implementation is missing (as with floats and structs). 

A file with testing is included.

Any PRs, advice, or help will be appreciated.
