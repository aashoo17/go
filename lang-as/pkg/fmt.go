package pkg

/*
Printf:
this is more or less works like as in c
I will explain basics one based on type
integer:
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

float:
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

bool:
%t	the word true or false

string:
%s	the uninterpreted bytes of the string or slice

slice:
%p	address of 0th element in base 16 notation, with leading 0x

pointer:
%p	base 16 notation, with leading 0x


Println:
println() is like printf() call but some default of each type is used
so in printf() we have multiple options but in println() one is used by go automatically
*/
