# CockroachDB - 2 Hour Workshop

In this workshop, we'll explore the following topics:

* Setting up multi-region CockroachDB Serverless clusters

* Connecting to CockroachDB from code

* Generating test data for your CockroachDB databases

* Using CockroachDB's Intelligent Insights to diagnose query performance bottlenecks

## Dependencies

To complete this workshop, you'll need to download the following dependencies:

* [CockroachDB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb)

* [dg](https://github.com/codingconcepts/dg/releases/latest)

``` sh
tar -xvf dg_[VERSION]_[OS].tar.gz
```

* [serve](https://github.com/codingconcepts/serve/releases/latest)

``` sh
tar -xvf serve_[VERSION]_[OS].tar.gz
```

## Workshop parts

#### CockroachDB Serverless

TODO(mb)

#### Connecting from Code

The resources/code_examples directory contains a number of examples for connecting to CockroachDB via your programming language of choice.

**/java_example**

The Java example is a very simple application that connects to a CockroachDB database. It makes use of [maven](https://maven.apache.org), so install that unless you'd prefer another tool.

Build the project

``` sh
cd resources/code_examples/java_example
mvn package
```

Run the project, substituting the value for the `CONNECTION_STRING` environmant variable as required.

``` sh
CONNECTION_STRING="jdbc:postgresql://localhost:26257/test?user=root" \
  java -jar target/hello-cockroach-0.1.0.jar
```

**/go_example**

The Go example is a very simple application that connects to a CockroachDB database. It uses Go's built-in `go mod` package manager, so no additional build dependencies are required.

Run the project, substituting the value for the `CONNECTION_STRING` environment variable as required.

``` sh
cd resources/code_examples/go_example

CONNECTION_STRING="postgres://root@localhost:26257/test?sslmode=disable" \
  go run main.go
```