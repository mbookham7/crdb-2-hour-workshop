# Setting Up a Serverless Cluster

During this section of the lab you will learn how to use Cockroach Cloud to create a CockroachDB Serverless cluster. Once this is up and running, which only takes a few seconds, you will then learn how to connect to this with code and the CLI client.

To kick thinks off lets create our serverless cluster

## Setup Your Serverless Cluster

In your browser access the following URL

```bash
https://cockroachlabs.cloud/
```

Login to the console, in the centre of the screen click `Create Cluster`

![create-cluster](/images/serverless-setup/create-cluster.png)

Once you have clicked on `Create Cluster` you will be presented with the a wizard that walks you through the setup of your serverless cluster.

In this scenario we will be creating we will be creating a multi regional serverless cluster in AWS. Under the Cloud Provider pick AWS.

![select-your-cloud](/images/serverless-setup/select-your-cloud.png)

Next, we are going to select the regions. As we are going to deploy a multi-region serverless cluster we will need to add two additional regions. Please select the following regions for this scenario.

- N. Virginia (us-east-1)
- Oregon (us-west-2)
- Frankfurt (eu-central-1)

![select-your-regions](/images/serverless-setup/select-your-regions.png)

Your first cluster can be free if you stay with in the limits which we will within this lab. You must enter a unique name for your cluster now user you name to make it unique.

`yourname-serverless-lab`

Enter in the box below.

![name-your-cluster](/images/serverless-setup/name-your-cluster.png)

Once you have made your selection review them and hit create!! WE ARE GOING GLOBAL!! Once your cluster is created then you will need to create an SQL User. The user name will be based on the account you logged in with but the password will be randomly generated, you need to copy this to you notepad and keep this safe for further steps in the lab.

![create-sql-user](/images/serverless-setup/create-sql-user.png)

## Connecting to your Cluster via Code

The resources/code_examples directory contains a number of examples for connecting to CockroachDB via your programming language of choice. Pick any of the following languages to build your example, you don't have to do them all!

git clone this repo.
```git clone https://github.com/mbookham7/crdb-2-hour-workshop.git```

**/java_example**

If you don't have Java installed, visit the [downloads](https://www.oracle.com/uk/java/technologies/downloads) site and install it from there.

The Java example is a very simple application that connects to a CockroachDB database. It makes use of [maven](https://maven.apache.org).

Build the project

``` sh
cd resources/code_examples/java_example
mvn package
```

Run the project, substituting the value for the `CONNECTION_STRING` environmant variable as required.The `CONNECTION_STRING` can be found in the Cockroach Cloud UI. If you click on `Connect` in the top right hand corner. Then change the `Select option/language` to `Connection String` This will be displayed in the window below. Copy this and paste it in to the connection string below.

``` sh
CONNECTION_STRING="<serverless-connection-string>" \
  java -jar target/hello-cockroach-0.1.0.jar
```

**/go_example**

If you don't have Go installed, visit the [downloads](https://go.dev/dl) site and install it from there.

The Go example is a very simple application that connects to a CockroachDB database. It uses Go's built-in `go mod` package manager, so no additional dependencies are required.

Run the project, substituting the value for the `CONNECTION_STRING` environment variable as required. The `CONNECTION_STRING` can be found in the Cockroach Cloud UI. If you click on `Connect` in the top right hand corner. Then change the `Select option/language` to `Connection String` This will be displayed in the window below. Copy this and paste it in to the connection string below.

``` sh
cd resources/code_examples/go_example

CONNECTION_STRING="<serverless-connection-string>" \
  go run main.go
```

**/dotnet_core_example**

If you don't have dotnet core installed, visit the [downloads](https://dotnet.microsoft.com/en-us/download/dotnet/3.1) site and install it from there.

The dotnet example is a very simple application that connects to a CockroachDB database. It uses `nuget` for package management, so no additional dependencies are required.

Run the project, substituting the value for the `CONNECTION_STRING` environment variable as required. The `CONNECTION_STRING` can be found in the Cockroach Cloud UI. If you click on `Connect` in the top right hand corner. Then change the `Select option/language` to `Connection String` This will be displayed in the window below. Copy this and paste it in to the connection string below.



``` sh
cd resources/code_examples/dotnet_core_example

CONNECTION_STRING="<serverless-connection-string>" \
  dotnet run
```

## Connecting to your Cluster via CLI

To connect to your cluster you will first need the `CockraochDB Client`. You can do this from the console by clicking connect in the top right hand corner.

Below are the commands for Mac and Windows. Copy the correct one for your Operating System and run.

**MacOS**

```shell
curl https://binaries.cockroachdb.com/cockroach-v23.1.10.darwin-10.9-amd64.tgz | tar -xz; sudo cp -i cockroach-v23.1.10.darwin-10.9-amd64/cockroach /usr/local/bin/
```

**Windows Powershell**

```powershell
$ErrorActionPreference = "Stop"; [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; $null = New-Item -Type Directory -Force $env:appdata/cockroach; Invoke-WebRequest -Uri https://binaries.cockroachdb.com/cockroach-v23.1.10.windows-6.2-amd64.zip -OutFile cockroach.zip; Expand-Archive -Force -Path cockroach.zip; Copy-Item -Force cockroach/cockroach-v23.1.10.windows-6.2-amd64/cockroach.exe -Destination $env:appdata/cockroach; $Env:PATH += ";$env:appdata/cockroach"; # We recommend adding ";$env:appdata/cockroach" to the Path variable for your system environment. See https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_environment_variables#saving-changes-to-environment-variables for more information.
```

Once you have downloaded the CockroachDB Client you are now ready to connect to your cluster. From the same window, from the connect button in the top right hand corner you can copy your connection command. Click copy and paste the command into your terminal.

![connection-string](/images/serverless-setup/connection-string.png)


> Note that this connection string will use the CockroachDB global load balancer, ensuring you are connected to your closest region. If you wish to override this and provide a region manually, the connection string can be updated as follows:

```shell
cockroach sql --url "postgresql://<username>@<cluster-name>.<region>.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
```

Now we are connected to the cluster we are going to run some basic SQL commands to get familiar with CockroachDB.


## Querying your cluster

See which region you are currently connected to (note that CockroachDB might direct you to a region that's close to you but not specified in your region list):

Note that if you'd like to have more space in your terminal, you can enter the following command from the CockroachDB shell to shorten your prompt:

```sh
\set prompt1 %/>
```

Show the region your client is currently connected to.
```sql
SELECT gateway_region();
```

Now show the regions in your cluster, what can you see?
Which is the primary region?

```sql
SHOW regions;
```

Create a database (this can then be used instead of "defaultdb" in your connection string for subsequent connections):

```sql
CREATE DATABASE workshop
  PRIMARY REGION "aws-eu-central-1"
  REGIONS "aws-us-east-1", "aws-us-west-2";
```

Now select the new database.
```sql 
USE workshop;
```

Perform some basic CRUD operations. Create a table with a colum for ID and a value:

```sql
CREATE TABLE example (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "value" STRING NOT NULL
);
```

We are now going to insert three rows in to our new table.
```sql
INSERT INTO example ("value") VALUES
  ('a'), ('b'), ('c');
```

Now if we do a `SELECT` on our table we can see that CockroachDB has automatically generated UUIDs for each row.

```sql
SELECT * FROM example;
```

Example Output:
```workshop> SELECT * FROM example;                                                                                                                                                                                        
                   id                  | value
---------------------------------------+--------
  323a533b-41ae-464a-a66f-f3a4a05f5eda | d
  c53949c5-5acf-4a56-a2d5-ee0d5b13870e | b
  cf6185de-b08c-48c6-a979-8f57e1026a2b | c
  e42a82bf-10ec-4870-87b5-7e8bf677f4b3 | a
(4 rows)
```

Now we are going to `INSERT` another row into out table. We are doing this so that find it earlier later to delete it!

```sql
INSERT INTO example ("id", "value") VALUES
  ('323a533b-41ae-464a-a66f-f3a4a05f5eda', 'd');
```

Do another `SELECT` for the example table.
```sql
SELECT * FROM example;
```

Now do another `SELECT` from the example table where the id matches the id we added in the earlier step.
```sql
SELECT * FROM example WHERE id = '323a533b-41ae-464a-a66f-f3a4a05f5eda';
```

Now we are going to delete the row.
```sql
DELETE FROM example WHERE id = '323a533b-41ae-464a-a66f-f3a4a05f5eda';
```

Do a final `SELECT` for the example table. You will see our row is deleted!!
```sql
SELECT * FROM example;
```

As you can see CockroachDB behaves like typical relational database, now lets look at query performance.

[next](/scenarios/query_performance/README.md)