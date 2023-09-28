# Query Performance

In this scenario, CockroachDB will help us find and fix a query performance issue.

To follow this example, you'll need:

* [CockroachDB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb)

* [dg](https://github.com/codingconcepts/dg/releases/latest)

* [serve](https://github.com/codingconcepts/serve/releases/latest)

## Steps

#### Part 1

Create a local instance of CockroachDB:

``` sh
cockroach demo --insecure --no-example-database --max-sql-memory 1Gi
```

Create a table:

``` sql
CREATE TABLE member (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "full_name" STRING NOT NULL,
  "contact" JSON NOT NULL
);
```

Insert some data:

``` sql
INSERT INTO member ("id", "full_name", "contact") VALUES
  (
    '04925a14-fdd4-4261-bd08-c6575ce59dcb',
    'Rob Reid',
    '{ "email": "rob@example.com", "street": "1 Charmingly Befuddled Road, Englandshire", "city": "London", "zip": "EC1 ABC" }'
  );
```

Select the data out using a JSON query. In this query, we're asking for members with a zip code of "EC1 ABC":

``` sql
SELECT m.id, m.full_name
FROM member m
WHERE m.contact @> '{"zip": "EC1 ABC"}';
```

So far so good! You'll seeing single-digit latency, so the query seems to be performing well.

#### Part 2

Now we'll insert a bunch of data, to highlight an issue with our table.

Generate a CSV containing 500,000 member rows using the following command:

``` sh
./dg -c dg.yaml -o csvs
```

Make the CSV file available to CockroachDB:

``` sh
./serve -d csvs -p 8000
```

Import the CSV file into CockroachDB:

``` sql
IMPORT INTO member (
	id, full_name, contact
)
CSV DATA (
    'http://localhost:8000/member.csv'
)
WITH skip='1', nullif = '', allow_quoted_null;
```

Now run the same query as before:

``` sql
SELECT m.id, m.full_name
FROM member m
WHERE m.contact @> '{"zip": "EC1 ABC"}';
``` 

Not so fast now! And the bigger our member table gets, the slower this query will become.

Visit CockroachDB's [Intelligent Insights page](http://127.0.0.1:8080/#/insights?tab=Schema+Insights&ascending=false&columnTitle=insights) in the console. If you don't see an entry on this page asking you to create an index, run the slow query a few more times. You should see the following:

![insights](/images/query-performance/insights.png)

Click the "Create Index" button, then run your query again.

To confirm that the index has been created, run the following query:

``` sql
SELECT create_statement FROM [SHOW CREATE TABLE member];
```

[next](/scenarios/geo-partitioning/README.md)