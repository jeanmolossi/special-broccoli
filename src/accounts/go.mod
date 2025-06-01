module github.com/jeanmolossi/special-broccoli/accounts

go 1.19

replace github.com/jeanmolossi/special-broccoli/common => ../common

require (
	github.com/aws/aws-lambda-go v1.36.0
	github.com/jeanmolossi/special-broccoli/common 761b1cdaed39
)
