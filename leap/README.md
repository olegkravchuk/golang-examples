Leap

Given a year, report if it is a leap year.

The tricky thing here is that a leap year in the Gregorian calendar occurs:

on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
For example, 1997 is not a leap year, but 1996 is. 1900 is not a leap year, but 2000 is.

If your language provides a method in the standard library that does this look-up, pretend it doesn't exist and implement it yourself.

Notes

Though our exercise adopts some very simple rules, there is more to learn!

For a delightful, four minute explanation of the whole leap year phenomenon, go watch this youtube video.

To run the tests run the command go test from within the exercise directory.

If the test suite contains benchmarks, you can run these with the -bench flag:

go test -bench .
For more detailed info about the Go track see the help page.

Source

JavaRanch Cattle Drive, exercise 3 http://www.javaranch.com/leap.jsp

Submitting Incomplete Problems

It's possible to submit an incomplete solution so you can see how others have completed the exercise.