module github.com/jeanmolossi/special-broccoli/accounts

go 1.19

replace github.com/jeanmolossi/special-broccoli/common => ../common

require (
	github.com/aws/aws-lambda-go v1.36.0
	github.com/jeanmolossi/special-broccoli/common v0.0.0-00010101000000-000000000000
)
